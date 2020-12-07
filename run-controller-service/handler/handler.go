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
		log.Fatalln("Error encountered while setting up connection with file builder service:", err)
		return nil
	}

	if riverRunnerServiceClient, err = getRiverRunnerServiceClient(); err != nil {
		log.Fatalln("Error encountered while setting up connection to river runner service:", err)
		return nil
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
		h.handleExecutorInstanceProvisioned(ctx, event)
	case events.FileAssemblySuccess:
		h.handleFileAssemblySuccess(ctx, event)
	case events.FileEvaluationFinished:
		h.handleFileEvaluationFinished(ctx, event)
	case events.ServiceError:
		h.handleServiceError(ctx, event)
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
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
		TestRunId:    eventData.TestRunSpec.Id,
	}
	createFileResp, err := h.FileService.CreateFile(ctx, &fileSpec)
	if err != nil {
		log.Printf("Error encountered while creating file for test run (id: %v):%v\n", eventData.TestRunSpec.Id, err.Error())
		h.changeTestRunState(ctx, eventData.TestRunSpec.Id, testRunStates.TestRunState.Error, []byte(err.Error()))
		return
	}
	stateChangeMetadata, _ := json.Marshal(&proto.FileSpec{
		Id:           createFileResp.File.Id,
		Name:         fileSpec.Name,
		Size:         fileSpec.Size,
		MaxChunkSize: fileSpec.MaxChunkSize,
	})
	h.changeTestRunState(ctx, createFileResp.File.TestRunId, testRunStates.TestRunState.FileUpload, stateChangeMetadata)
}

func (h EventHandler) changeTestRunState(ctx context.Context, testRunId uint32, newState testRunStates.TestRunStateType, stateMetadata []byte) {
	var newStateSpec testRunServiceProto.TestRunStateSpec = testRunServiceProto.TestRunStateSpec{
		TestRunId:     testRunId,
		State:         string(newState),
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
	// update test run state
	h.changeTestRunState(ctx, fileSpec.TestRunId, testRunStates.TestRunState.FileUploadDone, []byte{})

	// request file system provisioning to resource manager service
	var fileSystemSpec *resourceManagerProto.FileSystemSpec = &resourceManagerProto.FileSystemSpec{
		TestRunId:   fileSpec.TestRunId,
		SizeInBytes: fileSpec.Size,
	}

	if _, err := h.ResourceManagerService.ProvisionFileSystem(ctx, fileSystemSpec); err != nil {
		log.Println("Error calling ResourceManagerService.ProvisionFileSystem: ", err)
		h.changeTestRunState(ctx, fileSpec.TestRunId, testRunStates.TestRunState.Error, []byte(errors.FileSystemCreationError(fileSystemSpec, err).Error()))
		return
	}

	// update test run state to reflect that file system provisioning was started
	h.changeTestRunState(ctx, fileSpec.TestRunId, testRunStates.TestRunState.ProvisionFs, []byte{})
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
		log.Println("Error appending test run's user bytes to context:", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.TestRunUserBytesContextAppendError(eventData, err).Error()))
		return
	}

	// update test run state - mark file system provisioning as done
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.ProvisionFsDone, []byte{})

	// trigger tested file assembly
	fileAssembleRequest := &fileBuilderProto.FileAssembleRequest{TestRunId: eventData.TestRunId}
	if _, err := h.FileBuilderService.AssembleFile(ctx, fileAssembleRequest); err != nil {
		log.Println("Error calling FileBuilderService.AssembleFile: ", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.AssembleFileRequestError(fileAssembleRequest, err).Error()))
		return
	}

	// update test run state to mark that file assembly process started
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.AssembleFile, []byte{})
}

func (h EventHandler) handleExecutorInstanceProvisioned(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.ExecutorInstanceProvisionedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, err))
		return
	}

	// update test run state - mark executor instance provisioning as done
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.ProvisionExecutorInstanceDone, []byte{})

	// here file execution should be triggered
	// IGNORE IT FOR NOW - do this only if changing pipeline architecture

	// update test run state to mark that file execution process started
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Evaluating, []byte{})
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
		log.Println("Error appending test run's user bytes to context:", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.TestRunUserBytesContextAppendError(eventData, err).Error()))
		return
	}
	// update test run state - mark file assembly process as done
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.AssembleFileDone, []byte{})

	// trigger River execution
	runRequest := &riverRunnerProto.RunRequest{TestRunId: eventData.TestRunId}
	if _, err := h.RiverRunnerService.RunRiver(ctx, runRequest); err != nil {
		log.Println("Error calling RiverRunnerService.RunRiver: ", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.RiverRunRequestError(runRequest, err).Error()))
		return
	}

	// update test run state - mark run process started
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Evaluating, []byte{})
}

func (h EventHandler) handleFileEvaluationFinished(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.FileEvaluationFinishedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, err))
		return
	}

	// append wss event target bytes to context
	if err := h.appendTestRunUserBytesToContext(&ctx, eventData.TestRunId); err != nil {
		log.Println("Error appending test run's user bytes to context:", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.TestRunUserBytesContextAppendError(eventData, err).Error()))
		return
	}
	// TODO: add logic to mark test run as failed or succeeded
	// update test run state - mark it as finished or succeeded
	// HARDCODED succeeded for now
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Succeeded, []byte{})
}

func (h EventHandler) handleServiceError(ctx context.Context, event *proto.Event) {
	// unmarshal event specific data
	var eventData *proto.ServiceErrorEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, err))
		return
	}

	// append wss event target bytes to context
	if err := h.appendTestRunUserBytesToContext(&ctx, eventData.TestRunId); err != nil {
		log.Println("Error appending test run's user bytes to context:", err)
		h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, []byte(errors.TestRunUserBytesContextAppendError(eventData, err).Error()))
		return
	}

	// set run process error
	h.changeTestRunState(ctx, eventData.TestRunId, testRunStates.TestRunState.Error, event.Data)
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
