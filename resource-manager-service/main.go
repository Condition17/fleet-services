package main

import (
	"github.com/Condition17/resource-manager-service/config"
	"github.com/Condition17/resource-manager-service/database"
	"github.com/Condition17/resource-manager-service/handler"
	"github.com/Condition17/resource-manager-service/model"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	"github.com/Condition17/fleet-services/test-run-service/repository"
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
	db.AutoMigrate(&model.FileSystem{})

	pubsub := googlepubsub.NewBroker(googlepubsub.ProjectID(configs.GoogleProjectID))
	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Broker(pubsub),
		micro.Version("latest"),
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
