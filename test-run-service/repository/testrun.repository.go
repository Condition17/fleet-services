package repository

import (
	"github.com/Condition17/fleet-services/test-run-service/model"
	"gorm.io/gorm"
)

type TestRunRepository struct {
	Repository

	DB *gorm.DB
}

func (r *TestRunRepository) Create(testRun *model.TestRun) (*model.TestRun, error) {
	if err := r.DB.Create(testRun).Error; err != nil {
		return nil, err
	}
	return testRun, nil
}

func (r *TestRunRepository) GetAll(userId uint32) ([]*model.TestRun, error) {
	var testRuns []*model.TestRun
	if err := r.DB.Where("user_id = ?", userId).Find(&testRuns).Error; err != nil {
		return testRuns, err
	}
	return testRuns, nil
}

func (r *TestRunRepository) GetTestRun(userId uint32, testRunId uint32) (*model.TestRun, error) {
	var testRun model.TestRun
	if err := r.DB.First(&testRun, "user_id = ? AND id = ?", userId, testRunId).Error; err != nil {
		return nil, err
	}
	return &testRun, nil
}

func (r *TestRunRepository) Delete(userId uint32, testRunId uint32) error {
	if err := r.DB.Where("user_id = ? AND id = ?", userId, testRunId).Delete(&model.TestRun{}).Error; err != nil {
		return err
	}
	return nil
}
