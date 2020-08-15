package repository

import (
	"context"
	"fmt"
	"math"
	pb "upload/proto/upload"

	"github.com/gofrs/uuid"

	"github.com/gomodule/redigo/redis"
)

type File struct {
	ID                string `redis:"id"`
	Name              string `redis:"name"`
	Size              int64  `redis:"size"`
	MaxChunkSize      int64  `redis:"maxChunkSize"`
	ChunksStoresCount int64  `redis:"chunksStoresCount"`
}

type Chunk struct {
	Sha2     string `json:"sha2"`
	RefCount int64  `json:"ref_count"`
}

type Chunks []*Chunk

type Repository interface {
	Create(ctx context.Context, file *File) (*File, error)
	Read(ctx context.Context, id string) (*File, error)
}

type FileRepository struct {
	DB redis.Conn
}

type CacheEntityType string

const (
	fileEntity  CacheEntityType = "file"
	storeEntity CacheEntityType = "store"
)

// --- Marshal/Unmarshal protobuf structs

// Marshal/Unmarshall pb.File
func MarshalFile(file *pb.File) *File {
	return &File{
		ID:                file.Id,
		Name:              file.Name,
		Size:              file.Size,
		MaxChunkSize:      file.MaxChunkSize,
		ChunksStoresCount: file.ChunksStoresCount,
	}
}

func UnmarshalFile(file *File) *pb.File {
	return &pb.File{
		Id:                file.ID,
		Name:              file.Name,
		Size:              file.Size,
		MaxChunkSize:      file.MaxChunkSize,
		ChunksStoresCount: file.ChunksStoresCount,
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
		Sha2:     chunk.Sha2,
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

func (r *FileRepository) Read(ctx context.Context, id string) (*File, error) {
	values, err := redis.Values(r.DB.Do("HGETALL", composeFileKey(id)))
	if err != nil {
		return nil, err
	}

	var file File
	err = redis.ScanStruct(values, &file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) Create(ctx context.Context, file *File) (*File, error) {
	u, _ := uuid.NewV4()
	file.ID = u.String()
	key := fmt.Sprintf(composeFileKey(file.ID))
	maxStoreSize := int64(math.Pow(2, 32) - 1)

	// initialize needed chunk lists
	totalChunksCount := int64(math.Floor(float64(file.Size) / float64(file.MaxChunkSize)))
	neededStoresCount := int64(1)

	if totalChunksCount > maxStoreSize {
		neededStoresCount = int64(math.Floor(float64(totalChunksCount) / float64(maxStoreSize)))
	}

	// create Redis hash associated to file
	file.ChunksStoresCount = neededStoresCount
	if _, err := r.DB.Do("HSET", redis.Args{}.Add(key).AddFlat(file)...); err != nil {
		return nil, err
	}

	return file, nil
}

func composeFileKey(fileId string) string {
	return fmt.Sprintf("%s:%s", fileEntity, fileId)
}
