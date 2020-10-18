package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	testrun "test-run/proto/test-run"
)

type TestRun struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *TestRun) Call(ctx context.Context, req *testrun.Request, rsp *testrun.Response) error {
	log.Info("Received TestRun.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *TestRun) Stream(ctx context.Context, req *testrun.StreamingRequest, stream testrun.TestRun_StreamStream) error {
	log.Infof("Received TestRun.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&testrun.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *TestRun) PingPong(ctx context.Context, stream testrun.TestRun_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&testrun.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
