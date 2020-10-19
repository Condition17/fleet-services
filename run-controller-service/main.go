package main

import (
	"run-controller-service/handler"
	"run-controller-service/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	runcontrollerservice "run-controller-service/proto/run-controller-service"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.run-controller-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	runcontrollerservice.RegisterRunControllerServiceHandler(service.Server(), new(handler.RunControllerService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.run-controller-service", service.Server(), new(subscriber.RunControllerService))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
