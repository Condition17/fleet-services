package model

import (
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	userModels "github.com/Condition17/fleet-services/user-service/model"
	"gorm.io/gorm"
)

type TestRun struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint
	User   userModels.User
}

func MarshalTestRun(testRun *proto.TestRun) *TestRun {
	return &TestRun{
		Model: gorm.Model{ID: uint(testRun.Id)},
		Name:  testRun.Name,
	}
}

func UnmarshalTestRun(testRun *TestRun) *proto.TestRun {
	return &proto.TestRun{
		Id:   uint32(testRun.ID),
		Name: testRun.Name,
	}
}
