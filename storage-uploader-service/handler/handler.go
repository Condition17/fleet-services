package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"log"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	cloudstorage "github.com/Condition17/fleet-services/storage-uploader-service/storage"
	"github.com/micro/go-micro/v2"
)

type ChunkDataHandler struct {
	baseservice.BaseHandler
	StorageClient *cloudstorage.GcsClient
}

func (h ChunkDataHandler) HandleChunkDataMessage(chunkDataMessage *fileServiceProto.ChunkDataMessage) {
	if err := h.StorageClient.UploadChunk(chunkDataMessage.Sha2, string(chunkDataMessage.Data)); err != nil {
		log.Printf("[SERVICE ERROR]: Error encountered on upload (chunk sha2: %v): %v\n", chunkDataMessage.Sha2, err)
		h.SendServiceError(context.Background(), chunkDataMessage.TestRunId, err)
		return
	}

	// inform other services that the chunk was uploaded
	uploadedChunkData, _ := json.Marshal(chunkDataMessage)
	h.SendStorageUploadedChunkData(context.Background(), uploadedChunkData)
}

func NewHandler(service micro.Service) func(context.Context, *pubsub.Message) {
	var storageClient *cloudstorage.GcsClient

	// Initialize Cloud storage client
	storageClient, e := cloudstorage.InitClient()
	if e != nil {
		log.Fatal(e)
	}

	var handler = ChunkDataHandler{
		BaseHandler:   baseservice.NewBaseHandler(service),
		StorageClient: storageClient,
	}

	return func(c context.Context, msg *pubsub.Message) {
		var chunkDataMessage *fileServiceProto.ChunkDataMessage

		msg.Ack()
		fmt.Printf("Received message: '%s'\n", msg.Data)
		if err := json.Unmarshal(msg.Data, &chunkDataMessage); err != nil {
			fmt.Printf("Error encountered while unmarshalling message: %v", err)
			return
		}

		handler.HandleChunkDataMessage(chunkDataMessage)

	}
}
