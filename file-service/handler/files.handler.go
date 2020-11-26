package handler

import (
	"context"
	"encoding/json"

	pb "github.com/Condition17/fleet-services/file-service/proto/file-service"

	"github.com/Condition17/fleet-services/file-service/model"

	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/micro/go-micro/v2/errors"
)

func (h *Handler) CreateFile(ctx context.Context, req *pb.File, res *pb.Response) error {
	file, err := h.FileRepository.Create(context.Background(), model.MarshalFile(req))
	if err != nil {
		return err
	}
	res.File = model.UnmarshalFile(file)

	return nil
}

func (h *Handler) ReadFile(ctx context.Context, req *pb.File, res *pb.Response) error {
	file, err := h.FileRepository.Read(ctx, model.MarshalFile(req).ID)

	if file == nil || file.ID == "" {
		return errors.NotFound(h.Service.Name(), "File not found")
	}

	if err != nil {
		return err
	}

	res.File = model.UnmarshalFile(file)

	return nil
}

func (h Handler) HandleChunkStorageUploadSuccess(ctx context.Context, file *pb.File) error {
	uploadedChunksCount, err := h.FileRepository.IncrementUploadedChunksCount(ctx, file.Id)

	if err != nil {
		return err
	}

	eventData, _ := json.Marshal(
		&pb.FileChunkUploadedEventData{
			FileId:              file.Id,
			TotalChunksCount:    file.TotalChunksCount,
			UploadedChunksCount: uint64(uploadedChunksCount),
		})
	h.SendEventToWssQueue(ctx, "fileChunkUploaded", eventData)

	// Notify run controller that the file was uploaded
	if uint64(uploadedChunksCount) == file.TotalChunksCount {
		uploadedFileData, _ := json.Marshal(
			&runControllerProto.FileUploadedEventData{
				FileSpec: &runControllerProto.FileSpec{
					Id:           file.Id,
					Name:         file.Name,
					Size:         file.Size,
					MaxChunkSize: file.MaxChunkSize,
				},
			},
		)
		h.SendRunStateEvent(ctx, "file.uploaded", uploadedFileData)
	}

	return nil
}
