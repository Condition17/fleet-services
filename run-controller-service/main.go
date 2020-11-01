package main

import (
	"log"

	"github.com/Condition17/fleet-services/run-controller-service/config"
	handler "github.com/Condition17/fleet-services/run-controller-service/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/broker/googlepubsub/v2"
)

const topic string = "test-run-state"

func main() {
	configs := config.GetConfig()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.run-controller-service"),
		micro.Broker(googlepubsub.NewBroker(googlepubsub.ProjectID(configs.GoogleProjectID))),
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
	_, err := msgBroker.Subscribe(topic, handler.NewHandler(service))

	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
