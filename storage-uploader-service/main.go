package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"

	"github.com/Condition17/fleet-services/storage-uploader-service/handler"

	"github.com/Condition17/fleet-services/storage-uploader-service/config"
	"github.com/micro/go-micro/v2"
)

const gcsUploadSubscription string = "chunk-gcs-upload-subs"

func main() {
	configs := config.GetConfig()
	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Get PubSub client
	client, err := pubsub.NewClient(context.Background(), configs.GoogleProjectID)
	if err != nil {
		log.Fatalf("Failed to creacte PubSub client: %v", err)
	}

	// Subscribe to topic
	go func() {
		log.Printf("Subscribing to '%s'\n", gcsUploadSubscription)
		cctx, cancel := context.WithCancel(context.Background())
		err = client.Subscription(gcsUploadSubscription).Receive(cctx, handler.NewHandler(service))
		defer cancel()
	}()

	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
