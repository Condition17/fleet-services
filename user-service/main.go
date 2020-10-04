package main

import (
	"github.com/Condition17/fleet-services/user-service/handler"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	userservice "github.com/Condition17/fleet-services/user-service/proto/user-service"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.user-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	userservice.RegisterUserServiceHandler(service.Server(), new(handler.UserService))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
