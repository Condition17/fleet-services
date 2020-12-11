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
	alreadyCreatedForFile, err := redis.Bool(conn.Do("EXISTS", composeFileChunkBindingKey(spec.FileId, sha2)))
	if err != nil {
		return "", alreadyCreatedForFile, err
	}

	if alreadyCreatedForFile {
		// chunk already updated for the current file
		// ensure chunk is specified as part of the file
		if _, err := conn.Do("HSET", composeFileChunkStoreKey(spec.FileId, getStoreIndex(spec.Index)), spec.Index, sha2); err != nil {
			return "", true, err
		}
		return sha2, true, nil
	}

	// the chunk was not already uploaded for the current file
	// but we need to verify if the chunk was uploaded at all - DEDUPLICATION
	alreadyCreatedForAnotherFile, err := redis.Bool(conn.Do("EXISTS", hashKey))
	if err != nil {
		return "", false, err
	}

	conn.Send("MULTI")
	if !alreadyCreatedForAnotherFile {
		// create chunk entity
		var hashData *model.Chunk = &model.Chunk{Sha2: sha2, Size: int64(len(spec.Data))}
		conn.Send("HSET", redis.Args{}.Add(hashKey).AddFlat(hashData)...)
	}

	// create file-chunk binding
	conn.Send("SET", composeFileChunkBindingKey(spec.FileId, sha2), spec.Index)

	// add chunk as a part of the file
	conn.Send("HSET", composeFileChunkStoreKey(spec.FileId, getStoreIndex(spec.Index)), spec.Index, sha2)
	_, err = conn.Do("EXEC")

	if err != nil {
		return "", alreadyCreatedForAnotherFile, err
	}

	return sha2, alreadyCreatedForAnotherFile, nil
}

func (r *ChunkRepository) GetByIndexInFile(ctx context.Context, fileId string, index uint64) (*model.Chunk, error) {
	conn := r.DB.Get()
	defer conn.Close()

	// obtain chunk key
	chunkSha2, err := redis.String(conn.Do("HGET", fmt.Sprintf("%v:%v", composeFileKey(fileId), getStoreIndex(index)), fmt.Sprintf("%v", index)))
	if err != nil {
		return nil, err
	}
	// get all chunk details using its key
	return r.GetChunk(ctx, chunkSha2)
}

func (r *ChunkRepository) GetChunk(ctx context.Context, sha2 string) (*model.Chunk, error) {
	conn := r.DB.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", composeChunkKey(sha2)))

	if err != nil {
		return nil, err
	}

	var chunk model.Chunk
	err = redis.ScanStruct(values, &chunk)
	if err != nil {
		return nil, err
	}

	return &chunk, nil
}
