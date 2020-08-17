package repository

import (
	"context"
	"file-service/model"
	"fmt"
	"math"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
)

type FileRepository struct {
	Repository

	DB redis.Conn
}

func (r *FileRepository) Read(ctx context.Context, id string) (*model.File, error) {
	values, err := redis.Values(r.DB.Do("HGETALL", composeFileKey(id)))
	if len(values) == 0 {
		return nil, nil
	}

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
	u, _ := uuid.NewV4()
	file.ID = u.String()
	key := fmt.Sprintf(composeFileKey(file.ID))
	maxStoreSize := int64(math.Pow(2, 32) - 1)

	// initialize needed chunk lists
	totalChunksCount := int64(math.Floor(float64(file.Size) / float64(file.MaxChunkSize)))
	neededStoresCount := uint32(1)

	if totalChunksCount > maxStoreSize {
		neededStoresCount = uint32(math.Floor(float64(totalChunksCount) / float64(maxStoreSize)))
	}

	// create Redis hash associated to file
	file.ChunksStoresCount = neededStoresCount
	if _, err := r.DB.Do("HSET", redis.Args{}.Add(key).AddFlat(file)...); err != nil {
		return nil, err
	}

	return file, nil
}
