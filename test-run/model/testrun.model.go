package model

import (
	proto "github.com/Condition17/fleet-services/user-service/proto/test-run"
	"gorm.io/gorm"
)

type TestRun struct {
	gorm.Model
	Name     string `gorm:"not null"`
}

func MarshalUser(testRun *proto.TestRun) *User {
	return &TestRun{
		Name:     testRun.Name,
	}
}

func UnmarshalUser(user *TestRun) *proto.TestRun {
	return &proto.User{
		Name:     testRun.Name,
	}
}
