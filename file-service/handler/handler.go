package handler

import (
	"file-service/pubsub"
	"file-service/repository"
)

const gcsUploadTopic string = "chunk-gcs-upload"

type Service struct {
	Name            string
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
	MessagesBroker  pubsub.MessagesBroker
}
