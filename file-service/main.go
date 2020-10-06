package main

import (
	"github.com/Condition17/fleet-services/file-service/config"
	"github.com/Condition17/fleet-services/file-service/handler"
	"github.com/Condition17/fleet-services/file-service/pubsub"
	"github.com/Condition17/fleet-services/file-service/repository"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"
)

func main() {
	// Get configs
	config := config.GetConfig()

	// New Service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version("latest"),
	)

	// Setup Redis client
	redisPool := CreateRedisPool(config.RedisUrl)
	// ensure that connection to Redis is always properly closed

	// test redis connectivity via PING
	conn := redisPool.Get()
	if err := PingRedis(conn); err != nil {
		log.Fatal(err)
	} else {
		log.Info("Successfully connected to Redis")
	}
	conn.Close()

	// Initialise service
	service.Init()

	// Register Handler
	serviceHandler := handler.Service{
		Name:            config.ServiceName,
		FileRepository:  repository.FileRepository{DB: redisPool},
		ChunkRepository: repository.ChunkRepository{DB: redisPool},
		MessagesBroker:  pubsub.MessagesBroker{Broker: service.Server().Options().Broker},
	}
	if err := pb.RegisterFileServiceHandler(service.Server(), &serviceHandler); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
