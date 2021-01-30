package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Condition17/fleet-services/file-service/model"
	proto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/micro/go-micro/v2/metadata"
)

func (h Handler) HandleEvent(msg *pubsub.Message) {
	var message *proto.ChunkDataMessage

	if err := json.Unmarshal(msg.Data, &message); err != nil {
		log.Printf("Error encountered while unmarshaling chunk data message %v: %v\n", message, err)
		return
	}

	file, err := h.FileRepository.Read(context.Background(), message.FileId)
	if file == nil {
		h.SendServiceError(context.Background(), message.TestRunId, errors.New(fmt.Sprintf("file entity not found:%v", message.FileId)))
		return
	}

	if err != nil {
		h.SendServiceError(context.Background(), message.TestRunId, errors.New(fmt.Sprintf("error retrieving file entity (id: %v): %v\n", message.FileId, err)))
		return
	}

	ctx := metadata.Set(context.Background(), "Authorization", string(message.Authorization))
	if err := h.HandleChunkStorageUploadSuccess(ctx, model.UnmarshalFile(file)); err != nil {
		log.Printf("[SERVICE ERROR]: Error while handling chunk storage upload success: %v", err)
		h.SendServiceError(context.Background(), file.TestRunId, err)
		return
	}
}
