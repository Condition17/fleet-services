package main

import (
	"file-service/config"
	"file-service/handler"
	"file-service/repository"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	pb "file-service/proto/file-service"
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
	redisConn := redisPool.Get()

	// ensure that connection to Redis is always properly closed
	defer redisConn.Close()

	// test redis connectivity via PING
	if err := PingRedis(redisConn); err != nil {
		log.Fatal(err)
	} else {
		log.Info("Successfully connected to Redis")
	}
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
		FileRepository:  repository.FileRepository{DB: redisConn},
		ChunkRepository: repository.ChunkRepository{DB: redisConn},
	}
	if err := pb.RegisterFileServiceHandler(service.Server(), &serviceHandler); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
