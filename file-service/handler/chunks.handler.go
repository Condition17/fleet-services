package handler

import (
	"context"
	pb "file-service/proto/file-service"

	"github.com/micro/go-micro/errors"
)

func (s *Service) CreateChunk(ctx context.Context, req *pb.ChunkSpec, res *pb.EmptyResponse) error {
	if len(req.Data) == 0 {
		return errors.BadRequest(s.Name, "Could not create chunk: no chunk data transmitted.")
	}

	file, err := s.FileRepository.Read(ctx, req.FileId)

	if file == nil {
		return errors.BadRequest(s.Name, "File not found.")
	}

	if err != nil {
		return err
	}

	if req.Index >= file.TotalChunksCount {
		return errors.BadRequest(s.Name, "Invalid chunk index for file.")
	}

	if err := s.ChunkRepository.Create(ctx, req); err != nil {
		return err
	}

	return nil
}
