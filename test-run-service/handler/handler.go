package handler

import (
	"errors"
	"fmt"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/test-run-service/repository"
	googleStorage "github.com/Condition17/fleet-services/test-run-service/storage/google"
	"github.com/micro/go-micro/v2"
)

const runIssuesBucketName string = "fleet-runs-inputs"

type Handler struct {
	baseservice.BaseHandler
	TestRunRepository repository.TestRunRepository
	RunIssueRepository repository.RunIssueRepository
	cloudStorageClient *googleStorage.GcsClient
}

func NewHandler(service micro.Service, testRunRepo repository.TestRunRepository, runIssueRepo repository.RunIssueRepository) (*Handler, error) {
	var gcsClient *googleStorage.GcsClient
	var err error

	if gcsClient, err = googleStorage.GetClient(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error encountered while initializing google cloud storage client: %v", err))
	}

	return &Handler{
		BaseHandler:       baseservice.NewBaseHandler(service),
		TestRunRepository: testRunRepo,
		RunIssueRepository: runIssueRepo,
		cloudStorageClient: gcsClient,
	}, nil
}
