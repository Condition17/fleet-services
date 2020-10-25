package eventHandler

import (
	"context"
	"encoding/json"
	"log"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib"
	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	testRunServiceProto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
)

type handlerFunc = func(broker.Event) error

type eventHandler struct {
	FileService    fileServiceProto.FileService
	TestRunService testRunServiceProto.TestRunService
}

func (h eventHandler) buildCallContext(event *proto.Event) context.Context {
	return metadata.Set(context.Background(), "Token", string(event.Meta.Token))
}

func (h eventHandler) HandleTestRunCreated(event *proto.Event) {
	// unmarshal event speciffic data
	var eventData *proto.TestRunCreatedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Printf("Event (type: '%s') data unmarshall error: %v\n", event.Type, err)
		return
	}

	// create a new file for the received test run
	var fileSpec fileServiceProto.File = fileServiceProto.File{
		Name:         eventData.FileSpec.Name,
		Size:         eventData.FileSpec.Size,
		MaxChunkSize: eventData.FileSpec.MaxChunkSize,
	}
	createFileResp, err := h.FileService.CreateFile(h.buildCallContext(event), &fileSpec)
	if err != nil {
		// send a ws event in this case
		// include failure details
		log.Printf("File service error on create: %v", err)
		return
	}

	// assign the created file to the current test run
	var assignmentDetails testRunServiceProto.AssignRequest = testRunServiceProto.AssignRequest{
		TestRunId: eventData.TestRunSpec.Id,
		FileId:    createFileResp.File.Id,
	}
	if _, err := h.TestRunService.AssignFile(h.buildCallContext(event), &assignmentDetails); err != nil {
		// send a ws event in this case
		// include failure details
		log.Printf("Assign file error: %v", err)
		return
	}
}

func handleEvent(handler eventHandler, event *proto.Event) {
	switch event.Type {
	case events.TEST_RUN_CREATED:
		handler.HandleTestRunCreated(event)
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
}

func newEventHandler() eventHandler {
	return eventHandler{
		FileService:    fileServiceProto.NewFileService(lib.GetFullExternalServiceName("file-service"), client.DefaultClient),
		TestRunService: testRunServiceProto.NewTestRunService(lib.GetFullExternalServiceName("test-run-service"), client.DefaultClient),
	}
}

func New() handlerFunc {
	var handler eventHandler = newEventHandler()

	return func(e broker.Event) error {
		var event *proto.Event

		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			return err
		}
		handleEvent(handler, event)

		return nil
	}
}
