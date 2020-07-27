package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	upload "upload/proto/upload"
)

type Upload struct{}

func (e *Upload) Handle(ctx context.Context, msg *upload.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *upload.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
