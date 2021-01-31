package main

import (
	"context"
	"github.com/Condition17/fleet-services/storage-uploader-service/config"
	"github.com/Condition17/fleet-services/storage-uploader-service/handler"
	"github.com/micro/go-micro/v2"
	"log"
	"runtime"
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
	serviceHandler := handler.NewHandler(service)

	// Subscribe to topic
	go func() {
		log.Printf("Subscribing to '%s'\n", gcsUploadSubscription)
		sub := serviceHandler.PubSubClient.Subscription(gcsUploadSubscription)
		sub.ReceiveSettings.Synchronous = false
		sub.ReceiveSettings.NumGoroutines = runtime.NumCPU()
		cctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := sub.Receive(cctx, serviceHandler.GetPubSubMessageHandler())
		if err != nil {
			log.Fatalf("Subscribe error: %v", err)
		}
	}()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
