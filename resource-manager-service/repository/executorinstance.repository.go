package repository

import (
	"github.com/Condition17/fleet-services/resource-manager-service/model"
	"gorm.io/gorm"
)

type ExecutorInstanceRepository struct {
	Repository

	DB *gorm.DB
}

func (r *ExecutorInstanceRepository) Create(instance *model.ExecutorInstance) (*model.ExecutorInstance, error) {
	if err := r.DB.Create(instance).Error; err != nil {
		return nil, err
	}
	return instance, nil
}
