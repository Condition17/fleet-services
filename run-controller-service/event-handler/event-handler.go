package eventHandler

import (
	"context"
	"encoding/json"
	"log"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib"
	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
)

type eventHandlerFunc = func(broker.Event) error

func handleEvent(event *proto.Event) {
	switch event.Type {
	case events.TEST_RUN_CREATED:
		// should call create file
		var fileService fileServiceProto.FileService = fileServiceProto.NewFileService(lib.GetFullExternalServiceName("file-service"), client.DefaultClient)
		var ctx context.Context = context.WithValue(context.Background(), "Token", event.Meta.Token)

		res, err := fileService.CreateFile(ctx, &fileServiceProto.File{Name: "testFile", Size: 1000000000000, MaxChunkSize: 100})
		if err != nil {
			log.Printf("File service create call error: %v", err)
		} else {
			log.Printf("File service create file call response: %v", res)
		}

	// should associate fileId to the created test run
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
}

func New() eventHandlerFunc {
	return func(e broker.Event) error {
		var event *proto.Event

		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			return err
		}
		handleEvent(event)

		return nil
	}
}
