package handler

import (
	"context"
	"encoding/json"
	"log"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/run-controller-service/errors"
	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	testRunServiceProto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
)

type EventHandler struct {
	baseservice.BaseHandler
	FileService    fileServiceProto.FileService
	TestRunService testRunServiceProto.TestRunService
}

func NewHandler(service micro.Service) func(broker.Event) error {
	var handler EventHandler = EventHandler{
		BaseHandler:    baseservice.NewBaseHandler(service),
		FileService:    fileServiceProto.NewFileService(lib.GetFullExternalServiceName("file-service"), client.DefaultClient),
		TestRunService: testRunServiceProto.NewTestRunService(lib.GetFullExternalServiceName("test-run-service"), client.DefaultClient),
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
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
}

func (h EventHandler) sendErrorToWssQueue(ctx context.Context, err error) {
	h.SendEventToWssQueue(ctx, events.WSS_ERROR, []byte(err.Error()))
}

func (h EventHandler) handleTestRunCreated(ctx context.Context, event *proto.Event) {
	// unmarshal event speciffic data
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
		FileId:    assignmentDetails.FileId,
	})
	h.SendEventToWssQueue(ctx, events.WSS_FILE_ENTITY_CREATED, wssEventData)
}

func (h EventHandler) handleFileUploaded(ctx context.Context, event *proto.Event) {
	// unmarshal event speciffic data
	var eventData *proto.FileUploadedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event.Data, event))
		return
	}

	var fileId string = eventData.FileSpec.Id
	// get test run associated to the uploaded file
	var testRunDetailsResp *testRunServiceProto.TestRunDetails
	testRunDetailsResp, err := h.TestRunService.GetByFileId(ctx, &testRunServiceProto.FileSpec{Id: fileId})

	if err != nil {
		h.sendErrorToWssQueue(ctx, errors.TestRunRetrievalError(eventData, err.Error()))
		return
	}

	// send data to the client using WSS
	wssEventData, _ := json.Marshal(&proto.FileUploadCompletedEventData{
		TestRunId: testRunDetailsResp.TestRun.Id,
		FileId:    fileId,
	})
	h.SendEventToWssQueue(ctx, events.WSS_FILE_UPLOAD_COMPLETED, wssEventData)
}
