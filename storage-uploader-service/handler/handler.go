package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	cloudstorage "github.com/Condition17/fleet-services/storage-uploader-service/storage"
	"github.com/micro/go-micro/v2"
	"log"
)

type ChunkDataHandler struct {
	baseservice.BaseHandler
	StorageClient *cloudstorage.GcsClient
}

func (h ChunkDataHandler) HandleChunkDataMessage(chunkDataMessage *fileServiceProto.ChunkDataMessage) {
	if err := h.StorageClient.UploadChunk(chunkDataMessage.Sha2, chunkDataMessage.Data); err != nil {
		log.Printf("[SERVICE ERROR]: Error encountered on upload (chunk sha2: %v): %v\n", chunkDataMessage.Sha2, err)
		h.SendServiceError(context.Background(), chunkDataMessage.TestRunId, err)
		return
	}

	// inform other services that the chunk was uploaded
	uploadedChunkData, _ := json.Marshal(chunkDataMessage)
	h.SendStorageUploadedChunkData(context.Background(), uploadedChunkData)
}

func NewHandler(service micro.Service) ChunkDataHandler {
	var storageClient *cloudstorage.GcsClient

	// Initialize Cloud storage client
	storageClient, e := cloudstorage.InitClient()
	if e != nil {
		log.Fatal(e)
	}

	return ChunkDataHandler{
		BaseHandler:   baseservice.NewBaseHandler(service),
		StorageClient: storageClient,
	}
}

func (h ChunkDataHandler) GetPubSubMessageHandler() func(context.Context, *pubsub.Message) {
	return func(c context.Context, msg *pubsub.Message) {
		var chunkDataMessage *fileServiceProto.ChunkDataMessage

		msg.Ack()
		if err := json.Unmarshal(msg.Data, &chunkDataMessage); err != nil {
			fmt.Printf("Error encountered while unmarshalling message: %v\n", err)
			return
		}

		h.HandleChunkDataMessage(chunkDataMessage)
	}
}