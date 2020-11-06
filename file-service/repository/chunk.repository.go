package repository

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/Condition17/fleet-services/file-service/model"
	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"

	"github.com/gomodule/redigo/redis"
)

type ChunkRepository struct {
	Repository

	DB *redis.Pool
}

func (r *ChunkRepository) Create(ctx context.Context, spec *pb.ChunkSpec) (string, bool, error) {
	conn := r.DB.Get()
	defer conn.Close()

	var sha2 string = fmt.Sprintf("%x", sha256.Sum256(spec.Data))
	var hashKey string = composeChunkKey(sha2)

	// check if the chunk was already uploaded for the given file
	alreadyUploaded, err := redis.Bool(conn.Do("EXISTS", composeFileChunkBindingKey(spec.FileId, sha2)))
	if err != nil {
		return "", false, err
	}

	if alreadyUploaded {
		// chunk already updated for the current file - we no longer need to perform another operations
		return sha2, true, nil
	}

	// the chunk was not already uploaded for the current file
	// but we need to verify if the chunk was uploaded at all - DEDUPLICATION
	alreadyCreated, err := redis.Bool(conn.Do("EXISTS", hashKey))
	if err != nil {
		return "", false, err
	}

	conn.Send("MULTI")
	if !alreadyCreated {
		// create chunk
		var hashData *model.Chunk = &model.Chunk{Sha2: sha2, Size: int64(len(spec.Data))}
		conn.Send("HSET", redis.Args{}.Add(hashKey).AddFlat(hashData)...)
		// TODO: add data to action upload queue
	}

	// create file-chunk binding
	conn.Send("SET", composeFileChunkBindingKey(spec.FileId, sha2), "")

	// add chunk as a part of the file
	var storeIndex uint64 = spec.Index / maxChunkStoreSize
	conn.Send("HSET", composeFileChunkStoreKey(spec.FileId, storeIndex), spec.Index, sha2)
	_, err = conn.Do("EXEC")

	if err != nil {
		return "", false, err
	}

	return sha2, alreadyCreated, nil
}
