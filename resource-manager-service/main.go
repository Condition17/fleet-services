package main

import (
	"github.com/Condition17/fleet-services/resource-manager-service/config"
	"github.com/Condition17/fleet-services/resource-manager-service/handler"
	"github.com/Condition17/fleet-services/resource-manager-service/model"
	"github.com/Condition17/fleet-services/resource-manager-service/repository"
	"github.com/Condition17/fleet-services/resource-manager-service/storage/database"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
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
	db.AutoMigrate(&model.FileSystem{}, &model.ExecutorInstance{})

	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.NewHandler(service, repository.FileSystemRepository{DB: db}, repository.ExecutorInstanceRepository{DB: db})
	proto.RegisterResourceManagerServiceHandler(service.Server(), &serviceHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
