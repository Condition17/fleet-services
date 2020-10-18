package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Condition17/fleet-services/test-run-service/model"
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	userServiceProto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	microErrors "github.com/micro/go-micro/v2/errors"
	"gorm.io/gorm"
)

func getUserFromContex(ctx context.Context) (*userServiceProto.User, error) {
	usrBytes, _ := ctx.Value("User").([]byte)
	var user *userServiceProto.User
	if err := json.Unmarshal(usrBytes, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Create(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	user, _ := getUserFromContex(ctx)
	newTestRun := model.MarshalTestRun(req)
	newTestRun.UserID = user.Id
	createdTestRun, err := s.TestRunRepository.Create(newTestRun)
	if err != nil {
		return microErrors.BadRequest(s.Name, fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(createdTestRun)

	return nil
}

func (s *Service) Get(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	user, _ := getUserFromContex(ctx)
	result, err := s.TestRunRepository.GetTestRun(user.Id, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.Unauthorized(s.Name, "Test run with this id not found")
		}
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (s *Service) List(ctx context.Context, req *proto.EmptyRequest, res *proto.ListResponse) error {
	user, _ := getUserFromContex(ctx)
	results, err := s.TestRunRepository.GetAll(user.Id)
	if err != nil {
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}
	res.TestRuns = model.UnmarshalTestRunsCollection(results)

	return nil
}

func (s *Service) Delete(ctx context.Context, req *proto.TestRun, res *proto.EmptyResponse) error {
	user, _ := getUserFromContex(ctx)
	if err := s.TestRunRepository.Delete(user.Id, req.Id); err != nil {
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}

	return nil
}
