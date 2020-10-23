package handler

import (
	"context"
	"encoding/json"

	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"

	"github.com/micro/go-micro/errors"
)

type ChunkDataMessage struct {
	Sha2 string `json:"sha2"`
	Data []byte `json:"data"`
}

func (h *Handler) CreateChunk(ctx context.Context, req *pb.ChunkSpec, res *pb.EmptyResponse) error {
	if len(req.Data) == 0 {
		return errors.BadRequest(h.Service.Name(), "Could not create chunk: no chunk data transmitted.")
	}

	file, err := h.FileRepository.Read(ctx, req.FileId)
	if err != nil {
		return err
	}

	if file == nil {
		return errors.BadRequest(h.Service.Name(), "File not found.")
	}

	if req.Index >= file.TotalChunksCount {
		return errors.BadRequest(h.Service.Name(), "Invalid chunk index for file.")
	}

	chunkSHA2, err := h.ChunkRepository.Create(ctx, req)
	if err != nil {
		return err
	}

	if chunkSHA2 == "" {
		// This chunk has been already updated
		return nil
	}

	uploadData, _ := json.Marshal(&pb.ChunkDataMessage{Sha2: chunkSHA2, Data: req.Data})
	h.SendChunkDataToUploadQueue(ctx, uploadData)

	return nil
}
