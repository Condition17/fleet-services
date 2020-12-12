package repository

import (
	"github.com/Condition17/fleet-services/test-run-service/model"
	"github.com/Condition17/fleet-services/test-run-service/run-states"
	"gorm.io/gorm"
	"log"
)

type TestRunRepository struct {
	Repository

	DB *gorm.DB
}

func (r *TestRunRepository) Create(testRun *model.TestRun) (*model.TestRun, error) {
	testRun.State = runStates.TestRunState.Initiated
	if err := r.DB.Create(testRun).Error; err != nil {
		return nil, err
	}
	return testRun, nil
}

func (r *TestRunRepository) Update(newTestRun *model.TestRun) error {
	if err := r.DB.Save(newTestRun).Error; err != nil {
		return err
	}
	return nil
}

func (r *TestRunRepository) GetAll(userId uint32) ([]*model.TestRun, error) {
	var testRuns []*model.TestRun
	queryResult := r.DB.Table("test_runs").
		Select("test_runs.*, count(run_issues.id) as run_issues_count").
		Joins("left join run_issues on test_runs.id = run_issues.test_run_id").
		Where("test_runs.user_id = ?", userId).
		Group("test_runs.id").
		Find(&testRuns)

	if err := queryResult.Error; err != nil {
		log.Println("Error encountered:", err)
		return testRuns, err
	}

	return testRuns, nil
}

func (r *TestRunRepository) GetUserTestRun(userId uint32, testRunId uint32) (*model.TestRun, error) {
	var testRun model.TestRun
	if err := r.DB.Preload("RunIssues").First(&testRun, "user_id = ? AND id = ?", userId, testRunId).Error; err != nil {
		return nil, err
	}
	return &testRun, nil
}

func (r *TestRunRepository) GetTestRunById(testRunId uint32) (*model.TestRun, error) {
	var testRun model.TestRun

	if err := r.DB.Preload("RunIssues").Preload("User").First(&testRun, "id = ?", testRunId).Error; err != nil {
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
