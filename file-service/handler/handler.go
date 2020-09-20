package handler

import (
	"file-service/repository"
	"fmt"
	"log"

	"github.com/micro/go-micro/v2/broker"
)

type Service struct {
	Name            string
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
	PubSub          broker.Broker
}

func (s *Service) publishEvent() error {
	if err := s.PubSub.Publish("chunk-gcs-upload", &broker.Message{Body: []byte("This is a test message")}); err != nil {
		log.Printf("[pubsub] Publish failed: %v", err)
		return err
	}

	fmt.Println("[pubsub] Message successfully sent")
	return nil
}
