package handler

import (
	"file-service/repository"
)

type Service struct {
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
}
