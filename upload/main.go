package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"io"
	"os"
	"time"
	"upload/handler"

	pb "upload/proto/upload"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.upload"),
		micro.Version("latest"),
	)

	// Redis client
	rdb, err := CreateRedisConnection()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Successfully connected to Redis")
	}

	// --- TRYing redis query

	val, err := rdb.Get(context.Background(), "test").Result()
	if err != nil {
		log.Fatalf("Encountered error while querying Redis: %v", err)
	}
	fmt.Println("test => ", val)

	// ------------

	// trying gcs connection

	bucket := "fleet-chunks"
	object := "test-object"
	ctx := context.Background()

	// Creates a client.
	fmt.Println("Setup gcs")

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Open local file.
	f, err := os.Open("notes.txt")
	if err != nil {
		log.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}
	fmt.Fprintf(wc, "Blob %v uploaded.\n", object)

	// ------------

	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterUploadHandler(service.Server(), &handler.Service{})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
