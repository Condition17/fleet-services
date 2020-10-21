package main

import (
	eventHandler "github.com/Condition17/fleet-services/run-controller-service/event-handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

const topic string = "test-run-state"

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.run-controller-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Get the message broker instance
	msgBroker := service.Server().Options().Broker
	if err := msgBroker.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe run state topic
	_, err := msgBroker.Subscribe(topic, eventHandler.New())

	if err != nil {
		log.Fatal(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
