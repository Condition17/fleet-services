package model

import (
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/Condition17/fleet-services/test-run-service/run-states"
	userModels "github.com/Condition17/fleet-services/user-service/model"
	"gorm.io/gorm"
)

type TestRun struct {
	gorm.Model
	Name          string `gorm:"not null;type:varchar(100);default:null"`
	FileID        string `gorm:"type:varchar(100)"`
	UserID        uint32
	User          userModels.User
	State         runStates.TestRunStateType
	StateMetadata string `gorm:"type:text;default:''"`
}

func MarshalTestRun(testRun *proto.TestRun) *TestRun {
	return &TestRun{
		Model:  gorm.Model{ID: uint(testRun.Id)},
		Name:   testRun.Name,
		FileID: string(testRun.FileId),
	}
}

func UnmarshalTestRun(testRun *TestRun) *proto.TestRun {
	userData := userModels.UnmarshalUser(&testRun.User)
	return &proto.TestRun{
		Id:     uint32(testRun.ID),
		Name:   testRun.Name,
		FileId: string(testRun.FileID),
		UserId: testRun.UserID,
		User: &proto.User{
			Id:      userData.Id,
			Name:    userData.Name,
			Company: userData.Company,
			Email:   userData.Email,
		},
		State: string(testRun.State),
		StateMetadata: testRun.StateMetadata,
	}
}

func UnmarshalTestRunsCollection(testRuns []*TestRun) []*proto.TestRun {
	collection := make([]*proto.TestRun, 0)
	for _, run := range testRuns {
		collection = append(collection, UnmarshalTestRun(run))
	}
	return collection
}
