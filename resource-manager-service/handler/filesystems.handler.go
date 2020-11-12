package handler

import (
	"context"
	"fmt"
	"time"

	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	"google.golang.org/api/file/v1"
)

func (h *Handler) ProvisionFileSystem(ctx context.Context, req *proto.FileSystemSpec, res *proto.EmptyResponse) error {
	fileShareConfig := &file.FileShareConfig{CapacityGb: 1024, Name: "target"}
	networkConfig := &file.NetworkConfig{
		Modes:           []string{"MODE_IPV4"},
		Network:         "default",
		ReservedIpRange: "",
	}
	instance := &file.Instance{
		Description: "",
		FileShares:  []*file.FileShareConfig{fileShareConfig},
		Networks:    []*file.NetworkConfig{networkConfig},
		Tier:        "BASIC_HDD",
	}
	createCall := h.CloudFileStoreService.Projects.Locations.Instances.Create(
		fmt.Sprintf("projects/%s/locations/%s", "fleet-271114", "us-central1-a"), instance)

	createCall.InstanceId("tst123")
	operation, err := createCall.Do()

	if err != nil {
		fmt.Printf("Error encountered: %v", err)
	}

	fmt.Println(operation)

	h.waitForOperationToFinish(operation)

	return nil
}

func (h *Handler) waitForOperationToFinish(op *file.Operation) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			operationStatusCall := h.CloudFileStoreService.Projects.Locations.Operations.Get(op.Name)
			operation, err := operationStatusCall.Do()
			if err != nil {
				return fmt.Errorf("Locations.Operations.Get: %s", err)
			}

			if operation.Done == true {
				ticker.Stop()
				fmt.Printf("Response: %s", operation.Response)
				fmt.Printf("Metadata: %s", operation.Metadata)
				if operation.Error != nil {
					return fmt.Errorf("Operation %s failed with error: %v", op.Name, operation.Error)
				}
			}
		}
	}
}
