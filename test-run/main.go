package main

import (
	"test-run/handler"
	testrun "test-run/proto/test-run"

	"github.com/Condition17/fleet-services/common/auth"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// Get configs
	// Connect to database
	// db, err := CreateConnection()
	// if err != nil {
	// 	log.Fatalf("Could not connect to DB: %v", err)
	// }

	// defer db.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.test-run"),
		micro.Version("latest"),
		// auth middleware
		micro.WrapHandler(auth.ServiceAuthWrapper),
	)

	// Initialise service
	service.Init()

	// Register Handler
	testrun.RegisterTestRunHandler(service.Server(), new(handler.TestRun))

	// Create clients for another services
	// fileService := proto.NewFileService(common.GetFullExternalServiceName("file-service"), service.Client())
	// res, err := fileService.CreateFile(context.Background(), &proto.File{Name: "testFile", Size: 1000000000000, MaxChunkSize: 100})
	// if err != nil {
	// 	log.Fatalf("File service create call error: %v", err)
	// 	return
	// }

	// log.Infof("File service create call RESPONSE: %v", res)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
