package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	storageuploaderservice "storage-uploader-service/proto/storage-uploader-service"
)

type StorageUploaderService struct{}

func (e *StorageUploaderService) Handle(ctx context.Context, msg *storageuploaderservice.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *storageuploaderservice.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
