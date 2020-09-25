package main

import (
	"storage-uploader-service/handler"
	"storage-uploader-service/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	storageuploaderservice "storage-uploader-service/proto/storage-uploader-service"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.storage-uploader-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	storageuploaderservice.RegisterStorageUploaderServiceHandler(service.Server(), new(handler.StorageUploaderService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.storage-uploader-service", service.Server(), new(subscriber.StorageUploaderService))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
