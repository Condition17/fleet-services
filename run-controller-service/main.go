package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.run-controller-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
