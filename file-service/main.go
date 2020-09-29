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

	// --- TRYing redis query
	//
	//val, err := rdb.Get(context.Background(), "test").Result()
	//if err != nil {
	//	log.Fatalf("Encountered error while querying Redis: %v", err)
	//}
	//fmt.Println("test => ", val)

	// ------------

	// trying gcs connection
	//
	//bucket := "fleet-chunks"
	//object := "test-object"
	//ctx := context.Background()

	// Creates a client.
	//fmt.Println("Setup gcs")
	//
	//client, err := storage.NewClient(ctx)
	//if err != nil {
	//	log.Fatalf("Failed to create client: %v", err)
	//}
	//defer client.Close()
	//
	//// Open local file.
	//f, err := os.Open("notes.txt")
	//if err != nil {
	//	log.Errorf("os.Open: %v", err)
	//}
	//defer f.Close()
	//
	//ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	//defer cancel()
	//
	//// Upload an object with storage.Writer.
	//wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	//if _, err = io.Copy(wc, f); err != nil {
	//	log.Errorf("io.Copy: %v", err)
	//}
	//if err := wc.Close(); err != nil {
	//	log.Errorf("Writer.Close: %v", err)
	//}
	//fmt.Fprintf(wc, "Blob %v uploaded.\n", object)

	// ------------

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
