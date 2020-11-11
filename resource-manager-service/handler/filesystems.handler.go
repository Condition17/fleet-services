package handler

import (
	"context"
	"fmt"

	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
)

func (h *Handler) ProvisionFileSystem(ctx context.Context, req *proto.FileSystemSpec, res *proto.EmptyResponse) error {
	fmt.Println("ProvisionFileSystem")
	return nil
}
