package handler

import (
	"file-service/repository"
)

type Service struct {
	Name            string
	FileRepository  repository.FileRepository
	ChunkRepository repository.ChunkRepository
}
