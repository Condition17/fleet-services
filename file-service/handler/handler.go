package handler

import (
	"github.com/Condition17/fleet-services/file-service/repository"

	"github.com/Condition17/fleet-services/file-service/pubsub"
)

const gcsUploadTopic string = "chunk-gcs-upload"

type Service struct {
	Name            string
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
	MessagesBroker  pubsub.MessagesBroker
}
