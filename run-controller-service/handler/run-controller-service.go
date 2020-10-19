package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	runcontrollerservice "run-controller-service/proto/run-controller-service"
)

type RunControllerService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *RunControllerService) Call(ctx context.Context, req *runcontrollerservice.Request, rsp *runcontrollerservice.Response) error {
	log.Info("Received RunControllerService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *RunControllerService) Stream(ctx context.Context, req *runcontrollerservice.StreamingRequest, stream runcontrollerservice.RunControllerService_StreamStream) error {
	log.Infof("Received RunControllerService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&runcontrollerservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *RunControllerService) PingPong(ctx context.Context, stream runcontrollerservice.RunControllerService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&runcontrollerservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
