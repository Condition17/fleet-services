package main

import (
	"context"

	"github.com/Condition17/fleet-services/common"
	"github.com/Condition17/fleet-services/common/auth"
	"github.com/Condition17/fleet-services/test-run-service/config"
	"github.com/Condition17/fleet-services/test-run-service/handler"
	"github.com/Condition17/fleet-services/test-run-service/model"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/Condition17/fleet-services/test-run-service/storage/database"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
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
		// auth middleware
		micro.WrapHandler(auth.ServiceAuthWrapper),
	)

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.Service{
		Name:              configs.ServiceName,
		TestRunRepository: repository.TestRunRepository{DB: db},
	}
	proto.RegisterTestRunServiceHandler(service.Server(), &serviceHandler)

	fileService := fileServiceProto.NewFileService(common.GetFullExternalServiceName("file-service"), service.Client())
	res, err := fileService.CreateFile(context.Background(), &fileServiceProto.File{Name: "testFile", Size: 1000000000000, MaxChunkSize: 100})
	if err != nil {
		log.Fatalf("File service create call error: %v", err)
		return
	}

	log.Infof("File service create call RESPONSE: %v", res)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
