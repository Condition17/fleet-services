package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/Condition17/fleet-services/file-service/model"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
)

type FileRepository struct {
	Repository

	DB *redis.Pool
}

func (r *FileRepository) Read(ctx context.Context, id string) (*model.File, error) {
	conn := r.DB.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", composeFileKey(id)))

	if err != nil {
		return nil, err
	}

	var file model.File
	err = redis.ScanStruct(values, &file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) Create(ctx context.Context, file *model.File) (*model.File, error) {
	conn := r.DB.Get()
	defer conn.Close()

	u, _ := uuid.NewV4()
	file.ID = u.String()
	key := fmt.Sprintf(composeFileKey(file.ID))
	maxStoreSize := uint64(math.Pow(2, 32) - 1)

	// initialize needed chunk lists
	totalChunksCount := uint64(math.Round(float64(file.Size)/float64(file.MaxChunkSize) + 0.5))
	neededStoresCount := uint32(1)

	if totalChunksCount > maxStoreSize {
		neededStoresCount = uint32(math.Floor(float64(totalChunksCount) / float64(maxStoreSize)))
	}

	// create Redis hash associated to file
	file.ChunksStoresCount = neededStoresCount
	file.TotalChunksCount = totalChunksCount
	if _, err := conn.Do("HSET", redis.Args{}.Add(key).AddFlat(file)...); err != nil {
		return nil, err
	}

	return file, nil
}
