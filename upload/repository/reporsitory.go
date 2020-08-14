package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	pb "upload/proto/upload"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
)

type File struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Size              int64  `json:"size"`
	MaxChunkSize      int64  `json:"maxChunkSize"`
	ChunksStoresCount int64  `json:"chunksStoresCount"`
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
	DB *redis.Client
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

func (r *FileRepository) Read(ctx context.Context, id string) (*File, error) {
	fileKey := composeFileKey(id)
	res := r.DB.Get(context.Background(), fileKey)
	if err := res.Err(); err != nil {
		return nil, res.Err()
	}

	b, _ := res.Bytes()
	var file *File
	json.Unmarshal(b, &file)

	return file, nil
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
	if err := r.DB.Set(ctx, key, file, 0).Err(); err != nil {
		return nil, err
	}

	return file, nil
}

func composeFileKey(fileId string) string {
	return fmt.Sprintf("%s:%s", fileEntity, fileId)
}
