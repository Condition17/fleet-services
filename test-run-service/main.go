package main

import (
	"github.com/Condition17/fleet-services/test-run-service/config"
	"github.com/Condition17/fleet-services/test-run-service/handler"
	"github.com/Condition17/fleet-services/test-run-service/model"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/Condition17/fleet-services/test-run-service/storage/database"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
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

	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.Service{
		Name:              configs.ServiceName,
		TestRunRepository: repository.TestRunRepository{DB: db},
	}
	proto.RegisterTestRunServiceHandler(service.Server(), &serviceHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
