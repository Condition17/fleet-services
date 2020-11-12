package handler

import (
	"context"

	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/resource-manager-service/repository"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"google.golang.org/api/file/v1"
)

type Handler struct {
	baseservice.BaseHandler
	CloudFileStoreService *file.Service
	FileSystemRepository  repository.FileSystemRepository
}

func NewHandler(service micro.Service, fileSystemRepo repository.FileSystemRepository) Handler {
	fileStoreService, err := file.NewService(context.Background())

	if err != nil {
		log.Fatalf("Error encountered while setting up cloud file store service: %v", err)
	}

	return Handler{
		BaseHandler:           baseservice.NewBaseHandler(service),
		CloudFileStoreService: fileStoreService,
		FileSystemRepository:  fileSystemRepo,
	}
}
