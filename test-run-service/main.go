package main

import (
	"log"

	"github.com/Condition17/fleet-services/lib/auth"
	"github.com/Condition17/fleet-services/test-run-service/config"
	"github.com/Condition17/fleet-services/test-run-service/handler"
	"github.com/Condition17/fleet-services/test-run-service/model"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/Condition17/fleet-services/test-run-service/storage/database"

	"github.com/micro/go-micro/v2"

	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/micro/go-plugins/broker/googlepubsub/v2"
)

func main() {
	configs := config.GetConfig()
	// Create database connection
	db, err := database.CreateConnection()

	if err != nil {
		log.Fatalf("Error encountered while connectiong to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc.
	db.AutoMigrate(&model.TestRun{})

	pubsub := googlepubsub.NewBroker(googlepubsub.ProjectID(configs.GoogleProjectID))
	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Broker(pubsub),
		micro.Version("latest"),
		// auth middleware
		micro.WrapHandler(auth.ServiceAuthWrapper),
	)

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.NewHandler(service, repository.TestRunRepository{DB: db})
	proto.RegisterTestRunServiceHandler(service.Server(), &serviceHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
