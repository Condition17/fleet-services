package main

import (
	"encoding/json"
	"log"

	cloudstorage "storage-uploader-service/cloud-storage"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
)

const gcsUploadTopic string = "chunk-gcs-upload"

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.storage-uploader-service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Initialize Cloud storage client
	storageClient, e := cloudstorage.InitClient()
	if e != nil {
		log.Fatal("Cloud storage setup error!")
		log.Fatal(e)
	}

	// Get the broker instance
	msgBroker := service.Server().Options().Broker
	if err := msgBroker.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to topic on the broker
	_, err := msgBroker.Subscribe(gcsUploadTopic, func(p broker.Event) error {
		var message *fileServiceProto.ChunkDataMessage
		if err := json.Unmarshal(p.Message().Body, &message); err != nil {
			return err
		}

		if err := storageClient.UploadChunk(message.Sha2, message.Data); err != nil {
			log.Fatalf("Error encountered on upload (chunk %v): %v", message.Sha2, err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
