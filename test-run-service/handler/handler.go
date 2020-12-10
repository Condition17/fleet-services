package handler

import (
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	"github.com/micro/go-micro/v2"
)

type Handler struct {
	baseservice.BaseHandler
	TestRunRepository repository.TestRunRepository
	RunIssueRepository repository.RunIssueRepository
}

func NewHandler(service micro.Service, testRunRepo repository.TestRunRepository, runIssueRepo repository.RunIssueRepository) Handler {
	return Handler{
		BaseHandler:       baseservice.NewBaseHandler(service),
		TestRunRepository: testRunRepo,
		RunIssueRepository: runIssueRepo,
	}
}
