package handler

import (
	"context"
	"encoding/json"
	"github.com/Condition17/fleet-services/file-service/model"
	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib/auth"
	"github.com/micro/go-micro/v2/errors"
	"log"
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

	chunkSHA2, mayExistInStorage, err := h.ChunkRepository.Create(ctx, req)
	if err != nil {
		return err
	}

	go func() {
		if mayExistInStorage == true {
			if err := h.HandleChunkStorageUploadSuccess(ctx, model.UnmarshalFile(file)); err != nil {
				log.Printf("[SERVICE ERROR]: Error while handling chunk storage upload success: %v", err)
				h.SendServiceError(context.Background(), file.TestRunId, err)
			}

			return
		}

		uploadData, _ := json.Marshal(&pb.ChunkDataMessage{
			Sha2:          chunkSHA2,
			Data:          req.Data,
			FileId:        req.FileId,
			TestRunId:     file.TestRunId,
			Authorization: auth.GetAuthorizationBytesFromContext(ctx),
		})
		h.SendChunkDataToUploadQueue(ctx, uploadData)
	}()

	return nil
}

func (h *Handler) GetChunkDetailsByIndexInFile(ctx context.Context, req *pb.ChunkSpec, res *pb.ChunkDetails) error {
	if req.FileId == "" {
		return errors.BadRequest(h.Service.Name(), "Invalid request")
	}

	chunk, err := h.ChunkRepository.GetByIndexInFile(ctx, req.FileId, req.Index)
	if err != nil {
		return err
	}

	if chunk == nil {
		return errors.NotFound(h.Service.Name(), "Chunk not found.")
	}
	res.Chunk = model.UnmarshalChunk(chunk)

	return nil
}
