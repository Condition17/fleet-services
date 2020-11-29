package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Condition17/fleet-services/lib/auth"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/Condition17/fleet-services/test-run-service/model"
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	microErrors "github.com/micro/go-micro/v2/errors"
	"gorm.io/gorm"
)

func (h *Handler) Create(ctx context.Context, req *proto.CreateTestRunRequest, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	newTestRun := model.MarshalTestRun(req.TestRun)
	newTestRun.UserID = user.Id
	createdTestRun, err := h.TestRunRepository.Create(newTestRun)
	if err != nil {
		return microErrors.BadRequest(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(createdTestRun)

	// send test run created event
	eventData, _ := json.Marshal(
		&runControllerProto.TestRunCreatedEventData{
			TestRunSpec: &runControllerProto.TestRunSpec{
				Id:   uint32(createdTestRun.ID),
				Name: createdTestRun.Name,
			},
			FileSpec: &runControllerProto.FileSpec{
				Name:         req.FileSpec.Name,
				Size:         req.FileSpec.Size,
				MaxChunkSize: req.FileSpec.MaxChunkSize,
			},
		},
	)
	h.SendRunStateEvent(ctx, "test-run.created", eventData)

	return nil
}

func (h *Handler) Get(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	result, err := h.TestRunRepository.GetUserTestRun(user.Id, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "Test run with this id not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (h *Handler) GetByFileId(ctx context.Context, req *proto.FileSpec, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	result, err := h.TestRunRepository.GetTestRunByFileId(user.Id, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "No test run associated with the speciffied file was found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (h *Handler) GetById(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	var isServiceCaller bool = ctx.Value("serviceCaller").(bool)

	if !isServiceCaller {
		return microErrors.Unauthorized(h.Service.Name(), "Caller not authorized for this operation")
	}

	result, err := h.TestRunRepository.GetTestRunById(req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "No test run associated with the speciffied id was found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (h *Handler) List(ctx context.Context, req *proto.EmptyRequest, res *proto.ListResponse) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	results, err := h.TestRunRepository.GetAll(user.Id)
	if err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRuns = model.UnmarshalTestRunsCollection(results)

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *proto.TestRun, res *proto.EmptyResponse) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	if err := h.TestRunRepository.Delete(user.Id, req.Id); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return nil
}

func (h *Handler) AssignFile(ctx context.Context, req *proto.AssignRequest, res *proto.EmptyResponse) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	testRun, err := h.TestRunRepository.GetUserTestRun(user.Id, req.TestRunId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "Test run not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	if testRun.FileID != "" {
		return microErrors.Conflict(h.Service.Name(), "Test run already has a file assigned to it.")
	}

	testRun.FileID = req.FileId
	if err := h.TestRunRepository.Update(testRun); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return nil
}
