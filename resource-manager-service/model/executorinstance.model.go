package model

import (
	testRunModels "github.com/Condition17/fleet-services/test-run-service/model"
	"gorm.io/gorm"
)

type ExecutorInstance struct {
	gorm.Model

	IP          string `gorm:"not null;type:varchar(100);default:null"`
	Name        string `gorm:"not null;type:varchar(100);default:null"`
	MachineType string `gorm:"not nulll;"`
	TestRunID   uint32 `gorm:"unique"`
	TestRun     testRunModels.TestRun
}
