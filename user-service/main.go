package main

import (
	"log"

	"github.com/Condition17/fleet-services/user-service/auth"
	"github.com/Condition17/fleet-services/user-service/handler"
	"github.com/Condition17/fleet-services/user-service/model"
	"github.com/Condition17/fleet-services/user-service/repository"
	"github.com/Condition17/fleet-services/user-service/storage/database"

	"github.com/micro/go-micro/v2"

	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
)

func main() {
	// Create database connection
	db, err := database.CreateConnection()

	if err != nil {
		log.Fatalf("Error encountered while connectiong to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc.
	db.AutoMigrate(&model.User{})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.user-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.Service{
		Name:           "go.micro.api.user-service",
		UserRepository: repository.UserRepository{DB: db},
		TokenService:   &auth.TokenService{Issuer: "go.micro.api.user-service"},
	}
	proto.RegisterUserServiceHandler(service.Server(), &serviceHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
