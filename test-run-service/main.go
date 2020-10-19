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
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	userServiceProto "github.com/Condition17/fleet-services/user-service/proto/user-service"
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

	userServiceClient := userServiceProto.NewUserService(common.GetFullExternalServiceName("user-service"), client.DefaultClient)
	res, err := userServiceClient.Create(context.Background(), &userServiceProto.User{Email: "Asdfa@dev.test", Password: "asdafdsfas12"})
	if err != nil {
		log.Fatalf("User service create call error: %v", err)
		return
	}

	log.Infof("File service create call RESPONSE: %v", res)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
