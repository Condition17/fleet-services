package model

import (
	testRunModels "github.com/Condition17/fleet-services/test-run-service/model"
	"gorm.io/gorm"
)

type FileSystem struct {
	gorm.Model

	IP        string `gorm:"not null;type:varchar(100);default:null"`
	Name      string `gorm:"not null;type:varchar(100);default:null"`
	Capacity  uint64 `gorm:"not nulll"`
	TestRunID uint32
	TestRun   testRunModels.TestRun
}
