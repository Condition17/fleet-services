package model

import (
	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"
)

type File struct {
	ID                string `redis:"id"`
	Name              string `redis:"name"`
	Size              int64  `redis:"size"`
	MaxChunkSize      uint32 `redis:"maxChunkSize"`
	ChunksStoresCount uint32 `redis:"chunksStoresCount"`
	TotalChunksCount  uint64 `redis:"totalChunksCount"`
	TestRunId         uint32 `redis:"testRunId"`
}

func MarshalFile(file *pb.File) *File {
	return &File{
		ID:                file.Id,
		Name:              file.Name,
		Size:              file.Size,
		MaxChunkSize:      file.MaxChunkSize,
		ChunksStoresCount: file.ChunksStoresCount,
		TotalChunksCount:  file.TotalChunksCount,
		TestRunId:         file.TestRunId,
	}
}

func UnmarshalFile(file *File) *pb.File {
	return &pb.File{
		Id:                file.ID,
		Name:              file.Name,
		Size:              file.Size,
		MaxChunkSize:      file.MaxChunkSize,
		ChunksStoresCount: file.ChunksStoresCount,
		TotalChunksCount:  file.TotalChunksCount,
		TestRunId:         file.TestRunId,
	}
}
