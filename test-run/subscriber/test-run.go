package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	testrun "test-run/proto/test-run"
)

type TestRun struct{}

func (e *TestRun) Handle(ctx context.Context, msg *testrun.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *testrun.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
