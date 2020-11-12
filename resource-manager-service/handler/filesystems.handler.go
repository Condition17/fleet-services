package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Condition17/fleet-services/resource-manager-service/model"
	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	"google.golang.org/api/file/v1"
)

func (h *Handler) ProvisionFileSystem(ctx context.Context, req *proto.FileSystemSpec, res *proto.EmptyResponse) error {
	// TODO ===== REFACTOR THIS
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
	// ------------------

	operation, err := createCall.Do()
	if err != nil {
		return err
	}

	go h.executePostFSCreateOperationSteps(req.TestRunId, operation)

	return nil
}

func (h *Handler) executePostFSCreateOperationSteps(testRunId uint32, op *file.Operation) {
	finishedOperation, err := h.waitForOperationToFinish(op)
	if err != nil {
		fmt.Printf("Error encountered while waiting filestore create operation finish %v", err)
		return
	}

	var fsInstance file.Instance
	if err := json.Unmarshal(finishedOperation.Response, &fsInstance); err != nil {
		fmt.Printf("Error encountered while unmarshalling operation response: %v", err)
		return
	}

	_, err = h.FileSystemRepository.Create(&model.FileSystem{
		IP:                  fsInstance.Networks[0].IpAddresses[0],
		Name:                fsInstance.Name,
		FileShareCapacityGb: fsInstance.FileShares[0].CapacityGb,
		FileShareName:       fsInstance.FileShares[0].Name,
		TestRunID:           testRunId,
	})

	if err != nil {
		fmt.Printf("Error encountered on file system db entity creation: %v", err)
		return
	}

	fmt.Println("FS db entity created (for the record). Now we have to communicate it to the run-controller service")
	// send data to run controller service
}

func (h *Handler) waitForOperationToFinish(op *file.Operation) (*file.Operation, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		operationStatusCall := h.CloudFileStoreService.Projects.Locations.Operations.Get(op.Name)
		operation, err := operationStatusCall.Do()
		if err != nil {
			return nil, err
		}

		if !operation.Done {
			continue
		}

		// the operation has finished
		ticker.Stop()
		if operation.Error != nil {
			return nil, errors.New(fmt.Sprintf("Operation %s failed with error: %v", op.Name, operation.Error))
		}

		return operation, nil
	}
}
