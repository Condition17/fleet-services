package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/Condition17/fleet-services/file-service/repository"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
)

type Handler struct {
	baseservice.BaseHandler
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
}

func (h Handler) GetEventsHandler() func(context.Context, *pubsub.Message) {
	return func(c context.Context, msg *pubsub.Message) {
		msg.Ack()
		//log.Printf("Received message: '%s'\n", msg.Data)
		h.HandleEvent(msg)
	}
}
