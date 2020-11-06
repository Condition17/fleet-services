package handler

import (
	"github.com/Condition17/fleet-services/file-service/repository"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
)

const gcsUploadTopic string = "chunk-gcs-upload"

type Handler struct {
	baseservice.BaseHandler
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
}

func NewHandler(service micro.Service, fileRepo repository.FileRepository, chunksRepo repository.ChunkRepository) Handler {
	return Handler{
		BaseHandler:     baseservice.NewBaseHandler(service),
		FileRepository:  fileRepo,
		ChunkRepository: chunksRepo,
	}
}

func (h Handler) GetEventsHandler() func(broker.Event) error {
	return func(e broker.Event) error {
		h.HandleEvent(e)

		return nil
	}
}
