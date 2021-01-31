package main

import (
	"context"
	"github.com/Condition17/fleet-services/file-service/config"
	"github.com/Condition17/fleet-services/file-service/handler"
	"github.com/Condition17/fleet-services/file-service/repository"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"runtime"

	"github.com/micro/go-micro/v2"
	"log"

	proto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib/auth"
)

const storageUploadedChunksSubscription = "storage-uploaded-chunks-subs"

func main() {
	// Get configs
	config := config.GetConfig()

	// New Service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version("latest"),
		// auth middleware
		micro.WrapHandler(auth.ServiceAuthWrapper),
	)

	// Setup Redis client
	redisPool := CreateRedisPool(config.RedisUrl)
	// ensure that connection to Redis is always properly closed

	// test redis connectivity via PING
	conn := redisPool.Get()
	defer conn.Close()
	if err := PingRedis(conn); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to Redis")
	}

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.Handler{
		BaseHandler:     baseservice.NewBaseHandler(service),
		FileRepository:  repository.FileRepository{DB: redisPool},
		ChunkRepository: repository.ChunkRepository{DB: redisPool},
	}

	if err := proto.RegisterFileServiceHandler(service.Server(), &serviceHandler); err != nil {
		log.Fatal(err)
	}

	// Subscribe to storage uploaded chunks topic
	go func() {
		log.Printf("Subscribing to '%s'\n", storageUploadedChunksSubscription)
		sub := serviceHandler.PubSubClient.Subscription(storageUploadedChunksSubscription)
		sub.ReceiveSettings.Synchronous = false
		sub.ReceiveSettings.NumGoroutines = runtime.NumCPU()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := sub.Receive(ctx, serviceHandler.GetEventsHandler())

		if err != nil {
			log.Fatalf("Subscribe error: %v", err)
		}
	}()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
