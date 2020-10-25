package eventHandler

import (
	"encoding/json"
	"fmt"

	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/micro/go-micro/v2/broker"
)

type eventHandlerFunc = func(broker.Event) error

func handleEvent(event *proto.Event) {
	switch event.Type {
	case events.TEST_RUN_CREATED:
		// should call create file
		fmt.Printf("Handling event: %v", event)
		// Create clients for another services
	// fileService := proto.NewFileService(common.GetFullExternalServiceName("file-service"), service.Client())
	// res, err := fileService.CreateFile(context.Background(), &proto.File{Name: "testFile", Size: 1000000000000, MaxChunkSize: 100})
	// if err != nil {
	// 	log.Fatalf("File service create call error: %v", err)
	// 	return
	// }

	// should associate fileId to the created test run
	default:
		fmt.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
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
