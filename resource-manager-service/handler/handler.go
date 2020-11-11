package handler

import (
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/resource-manager-service/repository"
	"github.com/micro/go-micro/v2"
)

type Handler struct {
	baseservice.BaseHandler
}

func NewHandler(service micro.Service, repo repository.FileSystemRepository) Handler {
	return Handler{
		BaseHandler: baseservice.NewBaseHandler(service),
	}
}
