package model

import (
	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"
)

type Chunk struct {
	Sha2              string `redis:"sha2"`
	Size              int64  `redis:"size"`
	UploadedToStorage bool   `redis:"uploadedToStorage"`
}

func MarshalChunk(chunk *pb.Chunk) *Chunk {
	return &Chunk{
		Sha2: chunk.Sha2,
		Size: chunk.Size,
	}
}

func UnmarshalChunk(chunk *Chunk) *pb.Chunk {
	return &pb.Chunk{
		Sha2: chunk.Sha2,
		Size: chunk.Size,
	}
}
