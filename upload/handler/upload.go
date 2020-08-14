package handler

import (
	"context"
	"fmt"
	pb "upload/proto/upload"
	"upload/repository"
)

type Service struct {
	Repo repository.Repository
}

func (s *Service) Create(ctx context.Context, req *pb.File, res *pb.Response) error {
	fmt.Println(repository.MarshalFile(req))
	file, err := s.Repo.Create(context.Background(), repository.MarshalFile(req))
	if err != nil {
		return err
	}
	res.File = repository.UnmarshalFile(file)
	return nil
}

func (s *Service) Read(ctx context.Context, req *pb.File, res *pb.Response) error {
	file, err := s.Repo.Read(ctx, repository.MarshalFile(req).ID)
	if err != nil {
		return err
	}
	res.File = repository.UnmarshalFile(file)
	return nil
}
