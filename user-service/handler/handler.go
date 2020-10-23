package handler

import (
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/user-service/auth"
	"github.com/Condition17/fleet-services/user-service/repository"
	"github.com/micro/go-micro/v2"
)

type Handler struct {
	baseservice.BaseHandler
	UserRepository repository.UserRepository
	TokenService   auth.AuthChecker
}

func NewHandler(service micro.Service, userRepo repository.UserRepository, tokenService auth.AuthChecker) Handler {
	return Handler{
		BaseHandler:    baseservice.NewBaseHandler(service),
		UserRepository: userRepo,
		TokenService:   tokenService,
	}
}
