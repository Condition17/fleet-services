package handler

import (
	"context"
	"file-service/model"
	pb "file-service/proto/file-service"
	"fmt"

	"github.com/micro/go-micro/v2/errors"
)

func (s *Service) CreateFile(ctx context.Context, req *pb.File, res *pb.Response) error {
	file, err := s.FileRepository.Create(context.Background(), model.MarshalFile(req))
	if err != nil {
		return err
	}
	res.File = model.UnmarshalFile(file)

	return nil
}

func (s *Service) ReadFile(ctx context.Context, req *pb.File, res *pb.Response) error {
	fmt.Println("File handler - readFile execution")
	return errors.BadRequest(s.Name, "Ghinion fraiere")
	// file, err := s.FileRepository.Read(ctx, model.MarshalFile(req).ID)
	// if err != nil {
	// 	return err
	// }

	// res.File = model.UnmarshalFile(file)

	// return nil
}
