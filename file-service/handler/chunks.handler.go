package handler

import (
	"context"
	"encoding/json"
	pb "file-service/proto/file-service"

	"github.com/micro/go-micro/errors"
)

type ChunkDataMessage struct {
	Sha2 string `json:"sha2"`
	Data []byte `json:"data"`
}

func (s *Service) CreateChunk(ctx context.Context, req *pb.ChunkSpec, res *pb.EmptyResponse) error {
	if len(req.Data) == 0 {
		return errors.BadRequest(s.Name, "Could not create chunk: no chunk data transmitted.")
	}

	file, err := s.FileRepository.Read(ctx, req.FileId)
	if err != nil {
		return err
	}

	if file == nil {
		return errors.BadRequest(s.Name, "File not found.")
	}

	if req.Index >= file.TotalChunksCount {
		return errors.BadRequest(s.Name, "Invalid chunk index for file.")
	}

	chunkSHA2, err := s.ChunkRepository.Create(ctx, req)
	if err != nil {
		return err
	}

	message, _ := json.Marshal(&ChunkDataMessage{Sha2: chunkSHA2, Data: req.Data})
	if err := s.MessagesBroker.PublishEvent(gcsUploadTopic, message); err != nil {
		return err
	}

	return nil
}
