package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Condition17/fleet-services/common/auth"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/Condition17/fleet-services/test-run-service/model"
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	microErrors "github.com/micro/go-micro/v2/errors"
	"gorm.io/gorm"
)

func (h *Handler) Create(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromContext(ctx)
	newTestRun := model.MarshalTestRun(req)
	newTestRun.UserID = user.Id
	createdTestRun, err := h.TestRunRepository.Create(newTestRun)
	if err != nil {
		return microErrors.BadRequest(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(createdTestRun)

	// send test run created event
	eventData, _ := json.Marshal(&runControllerProto.TestRunCreatedEventData{Id: uint32(createdTestRun.ID), Name: createdTestRun.Name})
	h.sendRunStateEvent(ctx, "test-run.created", string(eventData))

	return nil
}

func (h *Handler) Get(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromContext(ctx)
	result, err := h.TestRunRepository.GetTestRun(user.Id, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.Unauthorized(h.Service.Name(), "Test run with this id not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (h *Handler) List(ctx context.Context, req *proto.EmptyRequest, res *proto.ListResponse) error {
	user, _ := auth.GetUserFromContext(ctx)
	results, err := h.TestRunRepository.GetAll(user.Id)
	if err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRuns = model.UnmarshalTestRunsCollection(results)

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *proto.TestRun, res *proto.EmptyResponse) error {
	user, _ := auth.GetUserFromContext(ctx)
	if err := h.TestRunRepository.Delete(user.Id, req.Id); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return nil
}
