package model

import (
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"gorm.io/gorm"
)

type RunIssue struct {
	gorm.Model
	BinaryPath  string `gorm:"not null;type:varchar(100);default:null"`
	Issue       string `gorm:"type:text"`
	InputBinUrl string `gorm:"type:text"`
	InputBytesPreview []byte
	TestRunID   uint32
	TestRun     TestRun
}

func MarshalRunIssue(runIssue *proto.RunIssue) *RunIssue {
	var inputBytesPreview []byte
	var inputBytes []byte = runIssue.InputBytes

	if len(inputBytes) > 4 {
		inputBytesPreview = inputBytes[:4]
	} else {
		inputBytesPreview = inputBytes
	}

	return &RunIssue{
		BinaryPath: runIssue.BinaryPath,
		Issue: runIssue.Issue,
		InputBytesPreview: inputBytesPreview,
		TestRunID: runIssue.TestRunId,
	}
}

func UnmarshalRunIssue(runIssue *RunIssue) *proto.RunIssue {
	return &proto.RunIssue{
		BinaryPath: runIssue.BinaryPath,
		Issue: runIssue.Issue,
		InputBytesPreview: runIssue.InputBytesPreview,
		InputBinUrl: runIssue.InputBinUrl,
	}
}

func UnmarshalRunIssuesCollection(runIssues []*RunIssue) []*proto.RunIssue {
	collection := make([]*proto.RunIssue, 0)
	for _, issue := range runIssues {
		collection = append(collection, UnmarshalRunIssue(issue))
	}
	return collection
}