package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	pb "upload/proto/upload"
)

type Service struct{}

func (s *Service) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {
	log.Info("Create logical file object")
	res.FileId = "alksndmo13243as.dmaskd"
	return nil
}

func (s *Service) Chunk(ctx context.Context, req *pb.ChunkRequest, res *pb.Response) error {
	log.Info("Upload chunk request")
	log.Info(req.Chunk)
	return nil
}


//
//// Call is a single request handler called via client.Call or the generated client code
//func (e *Upload) Call(ctx context.Context, req *upload.Request, rsp *upload.Response) error {
//	log.Info("Received Upload.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}
//
//// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *Upload) Stream(ctx context.Context, req *upload.StreamingRequest, stream upload.Upload_StreamStream) error {
//	log.Infof("Received Upload.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Infof("Responding: %d", i)
//		if err := stream.Send(&upload.StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *Upload) PingPong(ctx context.Context, stream upload.Upload_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Infof("Got ping %v", req.Stroke)
//		if err := stream.Send(&upload.Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}
