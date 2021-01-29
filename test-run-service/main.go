package main

import (
	"context"
	"github.com/Condition17/fleet-services/lib/auth"
	"github.com/micro/go-micro/v2/server"
	"log"

	"github.com/Condition17/fleet-services/test-run-service/config"
	"github.com/Condition17/fleet-services/test-run-service/handler"
	"github.com/Condition17/fleet-services/test-run-service/model"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/Condition17/fleet-services/test-run-service/storage/database"

	"github.com/micro/go-micro/v2"

	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/micro/go-plugins/broker/googlepubsub/v2"
)

func ServiceAuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		log.Println("Req method:", req.Method())
		// TestRunService.RegisterRunIssue is whitelisted
		// this route does not need authentication or any additional information added in context
		if req.Method() == "TestRunService.RegisterRunIssue" || req.Method() == "TestRunService.ForceStop"{
			log.Println("Skip authenticate request")
			return fn(ctx, req, resp)
		}

		return auth.AuthenticateRequest(fn, ctx, req, resp)
	}
}

func main() {
	var err error
	var serviceHandler *handler.Handler

	configs := config.GetConfig()
	// Create database connection
	db, err := database.CreateConnection()

	if err != nil {
		log.Fatalf("Error encountered while connectiong to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc.
	_ = db.AutoMigrate(&model.TestRun{}, &model.RunIssue{})

	pubsub := googlepubsub.NewBroker(googlepubsub.ProjectID(configs.GoogleProjectID))
	// New Service
	service := micro.NewService(
		micro.Name(configs.ServiceName),
		micro.Broker(pubsub),
		micro.Version("latest"),
		// auth middleware
		micro.WrapHandler(ServiceAuthWrapper),
	)

	// Initialise service
	service.Init()

	// Register Handler
	if serviceHandler, err = handler.NewHandler(service, repository.TestRunRepository{DB: db}, repository.RunIssueRepository{DB: db}); err != nil {
		log.Fatal(err)
	}

	proto.RegisterTestRunServiceHandler(service.Server(), serviceHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
