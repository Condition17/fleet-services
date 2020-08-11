package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	pb "upload/proto/upload"
)

type File struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	MaxChunkSize int32 `json:"maxChunkSize"`
	Chunks Chunks `json:"chunks"`
}

type Chunk struct {
	Sha2     string `json:"sha2"`
	RefCount int64  `json:"ref_count"`
}

type Chunks []*Chunk

type Repository interface {
	Create(ctx context.Context, file *File) (*File, error)
}

type FileRepository struct {
	DB *redis.Client
}

type CacheEntityType string

const (
	fileEntity CacheEntityType = "file"
	listEntity CacheEntityType = "list"
)
// --- Marshal/Unmarshal protobuf structs

// Marshal/Unmarshall pb.File
func MarshalFile(file *pb.File) *File {
	return &File{
		ID:     file.Id,
		Name:   file.Name,
		Size:   file.Size,
		MaxChunkSize: file.MaxChunkSize,
		Chunks: MarshalChunksCollection(file.Chunks),
	}
}

func UnmarshalFile(file *File) *pb.File {
	return &pb.File{
		Id: file.ID,
		Name: file.Name,
		Size: file.Size,
		MaxChunkSize: file.MaxChunkSize,
		Chunks: UnmarshalChunksCollection(file.Chunks),
	}
}

// Marshal/Unmarshall pb.Chunk

func MarshalChunk(chunk *pb.Chunk) *Chunk {
	return &Chunk{
		Sha2:     chunk.Sha2,
		RefCount: chunk.RefCount,
	}
}

func UnmarshalChunk(chunk *Chunk) *pb.Chunk {
	return &pb.Chunk{
		Sha2: chunk.Sha2,
		RefCount: chunk.RefCount,
	}
}

// Marshal/Unmarshal pb.Chunk collection

func MarshalChunksCollection(chunks []*pb.Chunk) []*Chunk {
	collection := make([]*Chunk, 0)
	for _, chunk := range chunks {
		collection = append(collection, MarshalChunk(chunk))
	}
	return collection
}

func UnmarshalChunksCollection(chunks []*Chunk) []*pb.Chunk {
	collection := make([]*pb.Chunk, 0)
	for _, chunk := range chunks {
		collection = append(collection, UnmarshalChunk(chunk))
	}
	return collection
}


// Marshal/Unmarshal of the previous object needed to store/retrieve them from Redis cache

// Binary Marshal/Unmarshal for File struct
func (e *File) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *File) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	return nil
}

func (r *FileRepository) Create(ctx context.Context, file *File) (*File, error) {
	u, _ := uuid.NewV4()
	file.ID = u.String()
	key := fmt.Sprintf("%s-%s", fileEntity, file.ID)

	
	if err := r.DB.HSet(ctx, key, file, 0).Err(); err != nil {
		return nil, err
	}

	return file, nil
}
