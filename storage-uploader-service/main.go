package main

import (
	"log"

	"github.com/Condition17/fleet-services/storage-uploader-service/handler"

	"github.com/Condition17/fleet-services/storage-uploader-service/config"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/broker/googlepubsub/v2"
)

const gcsUploadTopic string = "chunk-gcs-upload"

func main() {
	configs := config.GetConfig()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.storage-uploader-service"),
		micro.Broker(googlepubsub.NewBroker(googlepubsub.ProjectID(configs.GoogleProjectID))),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Get the broker instance
	msgBroker := service.Server().Options().Broker
	if err := msgBroker.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to topic on the broker
	if _, err := msgBroker.Subscribe(gcsUploadTopic, handler.NewHandler(service)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
