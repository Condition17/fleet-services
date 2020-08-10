package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"test-run/handler"
	testrun "test-run/proto/test-run"
	proto "github.com/Condition17/fleet-services/upload/proto/upload"
)

func main() {
	// Connect to database
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	defer db.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.test-run"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	testrun.RegisterTestRunHandler(service.Server(), new(handler.TestRun))

	// Create clients for another services
	uploadServiceClient := proto.NewUploadService("go.micro.api.upload", service.Client())
	res, err := uploadServiceClient.Create(context.Background(), &proto.CreateRequest{FileSize: 1000})
	if err != nil {
		log.Fatalf("Upload service create call error: %v", err)
		return
	}

	log.Infof("Upload service create call RESPONSE: %v", res)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
