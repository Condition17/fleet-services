package repository

import (
	"github.com/Condition17/fleet-services/test-run-service/model"
	"gorm.io/gorm"
)

type RunIssueRepository struct {
	Repository

	DB *gorm.DB
}

func (r *RunIssueRepository) Create(runIssue *model.RunIssue) error {
	if err := r.DB.Create(runIssue).Error; err != nil {
		return err
	}
	return nil
}