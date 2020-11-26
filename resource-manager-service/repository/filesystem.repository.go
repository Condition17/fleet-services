package repository

import (
	"github.com/Condition17/fleet-services/resource-manager-service/model"
	"gorm.io/gorm"
)

type FileSystemRepository struct {
	Repository

	DB *gorm.DB
}

func (r *FileSystemRepository) Create(fileSystem *model.FileSystem) (*model.FileSystem, error) {
	if err := r.DB.Create(fileSystem).Error; err != nil {
		return nil, err
	}
	return fileSystem, nil
}

func (r *FileSystemRepository) GetByTestRunId(testRunId uint32) (*model.FileSystem, error) {
	var fileSystem model.FileSystem

	if err := r.DB.Preload("TestRun.User").First(&fileSystem, "test_run_id = ?", testRunId).Error; err != nil {
		return nil, err
	}

	return &fileSystem, nil
}
