package repository

import (
	"context"
	"crypto/sha256"
	"file-service/model"
	pb "file-service/proto/file-service"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type ChunkRepository struct {
	Repository

	DB redis.Conn
}

func (r *ChunkRepository) Create(ctx context.Context, spec *pb.ChunkSpec) error {
	var sha2 string = fmt.Sprintf("%x", sha256.Sum256(spec.Data))
	var hashKey string = composeChunkKey(sha2)

	// check if the chunk was already uploaded for the given file
	alreadyUploaded, err := redis.Bool(r.DB.Do("EXISTS", composeFileChunkBindingKey(spec.FileId, sha2)))
	if err != nil {
		return err
	}

	if alreadyUploaded {
		// chunk already updated for the current file - we no longe need to perform another operations
		return nil
	}

	// the chunk was not already uploaded for the current file
	// but we need to verify if the chunk was uploaded at all - DEDUPLICATION
	alreadyCreated, err := redis.Bool(r.DB.Do("EXISTS", hashKey))
	if err != nil {
		return err
	}

	r.DB.Send("MULTI")
	if !alreadyCreated {
		// create chunk
		var hashData *model.Chunk = &model.Chunk{Sha2: sha2, Size: int64(len(spec.Data))}
		r.DB.Send("HSET", redis.Args{}.Add(hashKey).AddFlat(hashData)...)
		// TODO: add data to action upload queue
	}

	// create file-chunk binding
	r.DB.Send("SET", composeFileChunkBindingKey(spec.FileId, sha2), "")

	// add chunk as a part of the file
	var storeIndex uint64 = spec.Index / maxChunkStoreSize
	r.DB.Send("HSET", composeFileChunkStoreKey(spec.FileId, storeIndex), spec.Index, sha2)
	_, err = r.DB.Do("EXEC")

	return err
}
