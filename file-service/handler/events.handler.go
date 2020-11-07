package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Condition17/fleet-services/file-service/model"
	proto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/micro/go-micro/v2/broker"
)

func (h Handler) HandleEvent(e broker.Event) {
	var message *proto.ChunkDataMessage

	if err := json.Unmarshal(e.Message().Body, &message); err != nil {
		log.Printf("Error encountered while unmarshaling chunk data message %v\n", message)
		return
	}

	file, err := h.FileRepository.Read(context.Background(), message.FileId)
	if file == nil {
		log.Printf("File entity not found: %v\n", message.FileId)
		return
	}

	if err != nil {
		log.Printf("Error encountered while retrieving file entity: %v\n", err)
		return
	}

	if err := h.HandleChunkStorageUploadSuccess(context.WithValue(context.Background(), "Authorization", message.Authorization), model.UnmarshalFile(file)); err != nil {
		log.Printf("Error encountered: %v\n", err)
		return
	}
}
