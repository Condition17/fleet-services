package main

import (
	"context"
	"log"
	"runtime"

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
		sub := serviceHandler.PubSubClient.Subscription(testRunStateSubscription)
		sub.ReceiveSettings.Synchronous = false
		sub.ReceiveSettings.NumGoroutines = runtime.NumCPU()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := sub.Receive(ctx, serviceHandler.GetPubSubMessageHandler())

		if err != nil {
			log.Fatalf("Subscribe error: %v", err)
		}
	}()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
