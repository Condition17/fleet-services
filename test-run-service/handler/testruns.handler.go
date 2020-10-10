package handler

import (
	"context"
	"fmt"

	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
)

func (s *Service) Create(ctx context.Context, req *proto.TestRun, res *proto.TestRun) error {
	fmt.Println("Called create")
	return nil
}

func (s *Service) GetAll(ctx context.Context, req *proto.EmptyRequest, res *proto.GetAllResponse) error {
	fmt.Println("Get all called")
	return nil
}
