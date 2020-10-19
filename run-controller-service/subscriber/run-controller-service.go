package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	runcontrollerservice "run-controller-service/proto/run-controller-service"
)

type RunControllerService struct{}

func (e *RunControllerService) Handle(ctx context.Context, msg *runcontrollerservice.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *runcontrollerservice.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
