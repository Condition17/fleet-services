package handler

import (
	"github.com/Condition17/fleet-services/test-run-service/repository"
)

type Service struct {
	Name              string
	TestRunRepository repository.TestRunRepository
}
