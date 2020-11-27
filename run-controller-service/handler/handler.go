package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2/metadata"
	"log"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	resourceManagerProto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	"github.com/Condition17/fleet-services/run-controller-service/errors"
	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	testRunServiceProto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
)

type EventHandler struct {
	baseservice.BaseHandler
	FileService            fileServiceProto.FileService
	TestRunService         testRunServiceProto.TestRunService
	ResourceManagerService resourceManagerProto.ResourceManagerService
}

//conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
//if err != nil {
//log.Fatalf("Did not connect: %v", err)
//}
//fmt.Println("Connection:", conn)
//defer conn.Close()
//
//client := proto.NewBinaryBuilderClient(conn)
//if resp, err := client.Hello(context.Background(), &proto.EmptyMessage{}); err != nil {
//fmt.Println("Error:", err)
//return
//} else {
//fmt.Println("Call response:", resp)
//return
//}

func NewHandler(service micro.Service) func(broker.Event) error {
	var handler EventHandler = EventHandler{
		BaseHandler:            baseservice.NewBaseHandler(service),
		FileService:            fileServiceProto.NewFileService(lib.GetFullExternalServiceName("fileservice"), client.DefaultClient),
		TestRunService:         testRunServiceProto.NewTestRunService(lib.GetFullExternalServiceName("test-run-service"), client.DefaultClient),
		ResourceManagerService: resourceManagerProto.NewResourceManagerService(lib.GetFullExternalServiceName("ResourceManagerService"), client.DefaultClient),
	}

	return func(e broker.Event) error {
		var event *proto.Event

		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			return err
		}
		handler.HandleEvent(event)

		return nil
	}
}

func (h EventHandler) HandleEvent(event *proto.Event) {
	// TODO: this is odd. Find a better way to do this
	ctx := metadata.Set(context.Background(), "Authorization", string(event.Meta.Authorization))

	switch event.Type {
	case events.TEST_RUN_CREATED:
		h.handleTestRunCreated(ctx, event)
	case events.FILE_UPLOADED:
		h.handleFileUploaded(ctx, event)
	case events.FILE_SYSTEM_CREATED:
		h.handleFileSystemCreated(ctx, event)
	case events.EXECUTOR_INSTANCE_CREATED:
		h.handleExecutorInstanceCreated(ctx, event)
	case events.FILE_ASSEMBLY_SUCCEEDED:
		h.handleFileAssemblySuccess(ctx, event)
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
}

func (h EventHandler) sendErrorToWssQueue(ctx context.Context, err error) {
	h.SendEventToWssQueue(ctx, events.WSS_ERROR, []byte(err.Error()))
}

func (h EventHandler) handleTestRunCreated(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.TestRunCreatedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, event))
		return
	}

	// create a new file for the received test run
	var fileSpec fileServiceProto.File = fileServiceProto.File{
		Name:         eventData.FileSpec.Name,
		Size:         eventData.FileSpec.Size,
		MaxChunkSize: eventData.FileSpec.MaxChunkSize,
	}
	createFileResp, err := h.FileService.CreateFile(ctx, &fileSpec)
	if err != nil {
		h.sendErrorToWssQueue(ctx, errors.FileCreationError(eventData.TestRunSpec, err.Error()))
		return
	}

	// assign the created file to the current test run
	var assignmentDetails testRunServiceProto.AssignRequest = testRunServiceProto.AssignRequest{
		TestRunId: eventData.TestRunSpec.Id,
		FileId:    createFileResp.File.Id,
	}
	if _, err := h.TestRunService.AssignFile(ctx, &assignmentDetails); err != nil {
		h.sendErrorToWssQueue(ctx, errors.FileAssignError(eventData.TestRunSpec, err.Error()))
		return
	}

	// send informations to the client through WS service
	wssEventData, _ := json.Marshal(&proto.FileEntityCreatedEventData{
		TestRunId: assignmentDetails.TestRunId,
		FileSpec: &proto.FileSpec{
			Id:           assignmentDetails.FileId,
			Name:         fileSpec.Name,
			Size:         fileSpec.Size,
			MaxChunkSize: fileSpec.MaxChunkSize,
		},
	})
	h.SendEventToWssQueue(ctx, events.WSS_FILE_ENTITY_CREATED, wssEventData)
}

func (h EventHandler) handleFileUploaded(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.FileUploadedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, event))
		return
	}

	var fileSpec *proto.FileSpec = eventData.FileSpec

	// get test run associated to the uploaded file
	var testRunDetailsResp *testRunServiceProto.TestRunDetails
	testRunDetailsResp, err := h.TestRunService.GetByFileId(ctx, &testRunServiceProto.FileSpec{Id: fileSpec.Id})

	if err != nil {
		h.sendErrorToWssQueue(ctx, errors.TestRunRetrievalError(eventData, err.Error()))
		return
	}

	// send data to the client using WSS
	wssUploadCompletedEventData, _ := json.Marshal(&proto.FileUploadCompletedEventData{
		TestRunId: testRunDetailsResp.TestRun.Id,
		FileId:    fileSpec.Id,
	})
	h.SendEventToWssQueue(ctx, events.WSS_FILE_UPLOAD_COMPLETED, wssUploadCompletedEventData)

	// request file system provisioning to resource manager service
	var fileSystemSpec *resourceManagerProto.FileSystemSpec = &resourceManagerProto.FileSystemSpec{
		TestRunId:   testRunDetailsResp.TestRun.Id,
		SizeInBytes: fileSpec.Size,
	}

	if _, err := h.ResourceManagerService.ProvisionFileSystem(ctx, fileSystemSpec); err != nil {
		h.sendErrorToWssQueue(ctx, errors.FileSystemCreationError(fileSystemSpec, events.WSS_ERROR))
		return
	}

	// send file system creation start event to WSS
	wssEventData, _ := json.Marshal(&proto.FileSystemCreateEventData{TestRunId: testRunDetailsResp.TestRun.Id})
	h.SendEventToWssQueue(ctx, events.WSS_FILE_SYSTEM_CREATION_START, wssEventData)
}

func (h EventHandler) handleFileSystemCreated(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.FileSystemCreateEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, event))
		return
	}

	// append wss event target bytes to context
	if err := h.appendTestRunUserBytesToContext(&ctx, eventData.TestRunId); err != nil {
		log.Println(err)
		return
	}
	// send wss event
	wssEventData, _ := json.Marshal(&proto.FileSystemCreateEventData{TestRunId: eventData.TestRunId})
	h.SendEventToWssQueue(ctx, events.WSS_FILE_SYSTEM_CREATION_COMPLETED, wssEventData)
}

func (h EventHandler) handleExecutorInstanceCreated(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.ExecutorInstanceCreateEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, err))
		return
	}

	// append wss event target bytes to context
	if err := h.appendTestRunUserBytesToContext(&ctx, eventData.TestRunId); err != nil {
		log.Println(err)
		return
	}

	// send executor instance created event
	wssEventData, _ := json.Marshal(&proto.FileSystemCreateEventData{TestRunId: eventData.TestRunId})
	h.SendEventToWssQueue(ctx, events.WSS_EXECUTOR_INSTANCE_CREATION_COMPLETED, wssEventData)
}

func (h EventHandler) handleFileAssemblySuccess(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.FileAssemblySucceededEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, err))
		return
	}

	// append wss event target bytes to context
	if err := h.appendTestRunUserBytesToContext(&ctx, eventData.TestRunId); err != nil {
		log.Println(err)
		return
	}
	// send wss event
	log.Println("Sending wss event now. EventData:", eventData)
	h.SendEventToWssQueue(ctx, events.WSS_FILE_SUCCESSFULLY_ASSEMBLED, event.Data)
}

func (h EventHandler) appendTestRunUserBytesToContext(ctx *context.Context, testRunId uint32) error {
	testRunSpec := &testRunServiceProto.TestRun{Id: testRunId}
	testRunDetails, err := h.TestRunService.GetById(*ctx, testRunSpec)
	if err != nil {
		return errors.TestRunRetrievalError(testRunSpec, err.Error())
	}

	userDataBytes, err := json.Marshal(&testRunDetails.TestRun.User)
	*ctx = context.WithValue(*ctx, "User", userDataBytes)

	return nil
}
