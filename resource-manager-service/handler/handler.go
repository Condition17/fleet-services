package handler

import (
	"context"

	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/resource-manager-service/repository"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/file/v1"
)

type Handler struct {
	baseservice.BaseHandler
	CloudFileStoreService      *file.Service
	CloudComputeEngineService  *compute.Service
	FileSystemRepository       repository.FileSystemRepository
	ExecutorInstanceRepository repository.ExecutorInstanceRepository
}

func NewHandler(service micro.Service, fileSystemRepo repository.FileSystemRepository, executorInstanceRepo repository.ExecutorInstanceRepository) Handler {
	fileStoreService, err := file.NewService(context.Background())

	if err != nil {
		log.Fatalf("Error encountered while setting up cloud file store service: %v", err)
	}

	computeEngineService, err := compute.NewService(context.Background())

	if err != nil {
		log.Fatalf("Error encountered while setting up cloud compute engine service: %v", err)
	}

	return Handler{
		BaseHandler:                baseservice.NewBaseHandler(service),
		CloudFileStoreService:      fileStoreService,
		CloudComputeEngineService:  computeEngineService,
		FileSystemRepository:       fileSystemRepo,
		ExecutorInstanceRepository: executorInstanceRepo,
	}
}
