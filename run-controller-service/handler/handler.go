package handler

import (
	"context"
	"encoding/json"
	fileBuilderProto "github.com/Condition17/fleet-services/file-builder/proto/file-builder"
	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib"
	baseService "github.com/Condition17/fleet-services/lib/base-service"
	resourceManagerProto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	riverRunnerProto "github.com/Condition17/fleet-services/river-runner/proto/river-runner"
	"github.com/Condition17/fleet-services/run-controller-service/config"
	"github.com/Condition17/fleet-services/run-controller-service/errors"
	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	testRunServiceProto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	testRunStates "github.com/Condition17/fleet-services/test-run-service/run-states"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"google.golang.org/grpc"
	"log"
)

type EventHandler struct {
	baseService.BaseHandler
	FileService            fileServiceProto.FileService
	TestRunService         testRunServiceProto.TestRunService
	ResourceManagerService resourceManagerProto.ResourceManagerService
	FileBuilderService     fileBuilderProto.FileBuilderClient
	RiverRunnerService     riverRunnerProto.RiverRunnerClient
}

func NewHandler(service micro.Service) func(broker.Event) error {
	var err error
	var fileBuilderServiceClient fileBuilderProto.FileBuilderClient
	var riverRunnerServiceClient riverRunnerProto.RiverRunnerClient

	if fileBuilderServiceClient, err = getFileBuilderServiceClient(); err != nil {
		//log.Fatalln("Error encountered while setting up connection with file builder service:", err)
		//return nil
	}

	if riverRunnerServiceClient, err = getRiverRunnerServiceClient(); err != nil {
		//log.Fatalln("Error encountered while setting up connection to river runner service:", err)
		//return nil
	}

	var handler EventHandler = EventHandler{
		BaseHandler:            baseService.NewBaseHandler(service),
		FileService:            fileServiceProto.NewFileService(lib.GetFullExternalServiceName("fileService"), client.DefaultClient),
		TestRunService:         testRunServiceProto.NewTestRunService(lib.GetFullExternalServiceName("testRunService"), client.DefaultClient),
		ResourceManagerService: resourceManagerProto.NewResourceManagerService(lib.GetFullExternalServiceName("resourceManagerService"), client.DefaultClient),
		FileBuilderService:     fileBuilderServiceClient,
		RiverRunnerService:     riverRunnerServiceClient,
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

func getFileBuilderServiceClient() (fileBuilderProto.FileBuilderClient, error) {
	configs := config.GetConfig()
	conn, err := grpc.Dial(configs.FileBuilderServiceUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return fileBuilderProto.NewFileBuilderClient(conn), nil
}

func getRiverRunnerServiceClient() (riverRunnerProto.RiverRunnerClient, error) {
	configs := config.GetConfig()
	conn, err := grpc.Dial(configs.RiverRunnerServiceUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return riverRunnerProto.NewRiverRunnerClient(conn), nil
}

func (h EventHandler) HandleEvent(event *proto.Event) {
	// TODO: this is odd. Find a better way to do this
	ctx := metadata.Set(context.Background(), "Authorization", string(event.Meta.Authorization))

	switch event.Type {
	case events.TestRunInitiated:
		h.handleTestRunInitiated(ctx, event)
	case events.FileChunksUploaded:
		h.handleFileChunksUploaded(ctx, event)
	case events.FileSystemProvisioned:
		h.handleFileSystemProvisioned(ctx, event)
	case events.ExecutorInstanceProvisioned:
		h.handleExecutorInstanceCreated(ctx, event)
	case events.FILE_ASSEMBLY_SUCCEEDED:
		h.handleFileAssemblySuccess(ctx, event)
	case events.TEST_RUN_FINISHED:
		h.handleTestRunFinished(ctx, event)
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
}

func (h EventHandler) sendErrorToWssQueue(ctx context.Context, err error) {
	h.SendEventToWssQueue(ctx, events.WSS_ERROR, []byte(err.Error()))
}

func (h EventHandler) handleTestRunInitiated(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.TestRunInitiatedEventData
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

	stateChangeMetadata, _ := json.Marshal(&proto.FileSpec{
		Id:           assignmentDetails.FileId,
		Name:         fileSpec.Name,
		Size:         fileSpec.Size,
		MaxChunkSize: fileSpec.MaxChunkSize,
	})
	h.changeTestRunState(ctx, assignmentDetails.TestRunId, testRunStates.TestRunState.FileUpload, stateChangeMetadata)
}

func (h EventHandler) changeTestRunState(ctx context.Context, testRunId uint32, newState testRunStates.TestRunStateType, stateMetadata []byte) {
	var newStateSpec testRunServiceProto.TestRunStateSpec = testRunServiceProto.TestRunStateSpec{
		TestRunId: testRunId,
		State: string(newState),
		StateMetadata: string(stateMetadata),
	}
	if _, err := h.TestRunService.ChangeState(ctx, &newStateSpec); err != nil {
		if newState != testRunStates.TestRunState.Error {
			// Modify test run state if encountered an error while interacting with test run service
			// Do this only if the initial (un-transmitted) state as not an error state
			h.changeTestRunState(ctx, testRunId, testRunStates.TestRunState.Error, []byte(err.Error()))
		} else {
			log.Printf("Could not set error state for test run (id:%v)\n", testRunId)
		}
		return
	}

	// Send state change data to WSS
	newStateSpecBytes, _ := json.Marshal(&newStateSpec)
	h.SendEventToWssQueue(ctx, events.WssTestRunStateChanged, newStateSpecBytes)
}

func (h EventHandler) handleFileChunksUploaded(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.FileChunksUploadedEventData
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

	// update test run state
	h.changeTestRunState(ctx, testRunDetailsResp.TestRun.Id, testRunStates.TestRunState.FileUploadDone, []byte{})

	// request file system provisioning to resource manager service
	var fileSystemSpec *resourceManagerProto.FileSystemSpec = &resourceManagerProto.FileSystemSpec{
		TestRunId:   testRunDetailsResp.TestRun.Id,
		SizeInBytes: fileSpec.Size,
	}

	if _, err := h.ResourceManagerService.ProvisionFileSystem(ctx, fileSystemSpec); err != nil {
		log.Println("Error calling ResourceManagerService.ProvisionFileSystem: ", err)
		h.changeTestRunState(ctx, testRunDetailsResp.TestRun.Id, testRunStates.TestRunState.Error, []byte(errors.FileSystemCreationError(fileSystemSpec, err).Error()))
		return
	}

	// update test run state to reflect that file system provisioning was started
	h.changeTestRunState(ctx, testRunDetailsResp.TestRun.Id, testRunStates.TestRunState.ProvisionFs, []byte{})
}

func (h EventHandler) handleFileSystemProvisioned(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.FileSystemProvisionedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, event))
		return
	}

	// append wss event target bytes to context
	if err := h.appendTestRunUserBytesToContext(&ctx, eventData.TestRunId); err != nil {
		log.Println(err)
		return
	}

	// update test run state - mark file system provisioning as done
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.ProvisionFsDone, []byte{})

	// trigger tested file assembly
	fileAssembleRequest := &fileBuilderProto.FileAssembleRequest{TestRunId: eventData.TestRunId}
	if _, err := h.FileBuilderService.AssembleFile(ctx, fileAssembleRequest); err != nil {
		log.Println("Error calling FileBuilderService.AssembleFile: ", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.FileSystemCreationError(fileAssembleRequest, err).Error()))
		return
	}

	// update test run state to mark that file assembly process started
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.AssembleFile, []byte{})
}

func (h EventHandler) handleExecutorInstanceCreated(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.ExecutorInstanceProvisionedEventData
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
	wssEventData, _ := json.Marshal(&proto.ExecutorInstanceProvisionedEventData{TestRunId: eventData.TestRunId})
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
	h.SendEventToWssQueue(ctx, events.WSS_FILE_SUCCESSFULLY_ASSEMBLED, event.Data)

	// trigger river execution
	runRequest := &riverRunnerProto.RunRequest{TestRunId: eventData.TestRunId}
	if _, err := h.RiverRunnerService.RunRiver(ctx, runRequest); err != nil {
		h.sendErrorToWssQueue(ctx, errors.AssembleFileRequestError(runRequest, err.Error()))
		return
	}
}

func (h EventHandler) handleTestRunFinished(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.RiverRunFinishedEventData
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
	h.SendEventToWssQueue(ctx, events.WSS_TEST_RUN_FINSHED, event.Data)
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
