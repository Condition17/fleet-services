package handler

import (
	"github.com/Condition17/fleet-services/user-service/auth"
	"github.com/Condition17/fleet-services/user-service/repository"
)

type Service struct {
	Name           string
	UserRepository repository.UserRepository
	TokenService   auth.AuthChecker
}
