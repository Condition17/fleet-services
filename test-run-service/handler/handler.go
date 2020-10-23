package handler

import (
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/micro/go-micro/v2"
)

type Handler struct {
	BaseHandler
	TestRunRepository repository.TestRunRepository
}

func NewHandler(service micro.Service, repo repository.TestRunRepository) Handler {
	return Handler{
		BaseHandler:       NewBaseHandler(service),
		TestRunRepository: repo,
	}
}
