package handler

import (
	"context"
	"encoding/json"
	"log"

	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	cloudstorage "github.com/Condition17/fleet-services/storage-uploader-service/storage"
)

type ChunkDataHandler struct {
	baseservice.BaseHandler
	StorageClient *cloudstorage.GcsClient
}

func (h ChunkDataHandler) HandleChunkDataMessage(chunkDataMessage *fileServiceProto.ChunkDataMessage) {
	if err := h.StorageClient.UploadChunk(chunkDataMessage.Sha2, chunkDataMessage.Data); err != nil {
		log.Printf("Error encountered on upload (chunk %v): %v\n", chunkDataMessage.Sha2, err)
		return
	}

	log.Printf("-----> Uploaded chunk data: %v\n", chunkDataMessage)
	// inform other services that the chunk was uploaded
	uploadedChunkData, _ := json.Marshal(chunkDataMessage)
	h.SendStorageUploadedChunkData(context.Background(), uploadedChunkData)
}

func NewHandler(service micro.Service) func(broker.Event) error {
	var storageClient *cloudstorage.GcsClient

	// Initialize Cloud storage client
	storageClient, e := cloudstorage.InitClient()
	if e != nil {
		log.Fatal(e)
	}

	var handler ChunkDataHandler = ChunkDataHandler{
		BaseHandler:   baseservice.NewBaseHandler(service),
		StorageClient: storageClient,
	}

	return func(e broker.Event) error {
		var chunkDataMessage *fileServiceProto.ChunkDataMessage
		if err := json.Unmarshal(e.Message().Body, &chunkDataMessage); err != nil {
			return err
		}

		handler.HandleChunkDataMessage(chunkDataMessage)

		return nil
	}
}
