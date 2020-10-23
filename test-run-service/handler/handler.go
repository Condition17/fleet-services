package handler

import (
	baseservice "github.com/Condition17/fleet-services/common/base-service"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/micro/go-micro/v2"
)

type Handler struct {
	baseservice.BaseHandler
	TestRunRepository repository.TestRunRepository
}

func NewHandler(service micro.Service, repo repository.TestRunRepository) Handler {
	return Handler{
		BaseHandler:       baseservice.NewBaseHandler(service),
		TestRunRepository: repo,
	}
}
