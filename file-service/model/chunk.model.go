package model

import (
	pb "file-service/proto/file-service"
)

type Chunk struct {
	Sha2     string `json:"sha2"`
	RefCount uint32 `json:"ref_count"`
}

func MarshalChunk(chunk *pb.Chunk) *Chunk {
	return &Chunk{
		Sha2:     chunk.Sha2,
		RefCount: chunk.RefCount,
	}
}

func UnmarshalChunk(chunk *Chunk) *pb.Chunk {
	return &pb.Chunk{
		Sha2:     chunk.Sha2,
		RefCount: chunk.RefCount,
	}
}
