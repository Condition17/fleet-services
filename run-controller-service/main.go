package main

import (
	"context"
	"log"

	"github.com/Condition17/fleet-services/run-controller-service/config"
	handler "github.com/Condition17/fleet-services/run-controller-service/handler"
	"github.com/micro/go-micro/v2"
)

const testRunStateSubscription = "test-run-state-subs"

func main() {
	configs := config.GetConfig()
	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Version("latest"),
	)

	// Initialize service
	service.Init()

	serviceHandler := handler.NewHandler(service)
	// Subscribe run state topic
	go func() {
		log.Printf("Subscribing to '%s'\n", testRunStateSubscription)
		cctx, cancel := context.WithCancel(context.Background())
		err := serviceHandler.PubSubClient.Subscription(testRunStateSubscription).
			Receive(cctx, serviceHandler.GetPubSubMessageHandler())
		defer cancel()

		if err != nil {
			log.Fatalf("Subscribe error: %v", err)
		}
	}()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
