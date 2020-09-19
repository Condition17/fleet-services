package main

import (
	"context"
	"test-run/handler"
	testrun "test-run/proto/test-run"

	config "test-run/config"

	common "github.com/Condition17/fleet-services/common"
	proto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// Get configs
	config := config.GetConfig()
	// Connect to database
	// db, err := CreateConnection()
	// if err != nil {
	// 	log.Fatalf("Could not connect to DB: %v", err)
	// }

	// defer db.Close()

	// New Service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	testrun.RegisterTestRunHandler(service.Server(), new(handler.TestRun))

	// Create clients for another services
	fileServiceClient := proto.NewFileService(common.GetFullExternalServiceName("file-service"), service.Client())
	res, err := fileServiceClient.CreateFile(context.Background(), &proto.File{Name: "testFile", Size: 1000000000000, MaxChunkSize: 100})
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
