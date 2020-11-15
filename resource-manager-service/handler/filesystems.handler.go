package handler

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/Condition17/fleet-services/resource-manager-service/config"
	"github.com/Condition17/fleet-services/resource-manager-service/model"
	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"google.golang.org/api/file/v1"
)

const MIN_FILESTORE_CAPACITY_GB int64 = 1024
const MAX_FILESTORE_CAPACITY_GB int64 = 63900

func (h *Handler) ProvisionFileSystem(ctx context.Context, req *proto.FileSystemSpec, res *proto.EmptyResponse) error {
	var requestedSizeGb int64 = int64(math.Round(float64(req.SizeInBytes)/float64(math.Pow10(9)) + 0.5))
	var neededFsCapacityGb = int64(math.Max(float64(MIN_FILESTORE_CAPACITY_GB), float64(requestedSizeGb)))

	if neededFsCapacityGb > MAX_FILESTORE_CAPACITY_GB {
		return errors.New(fmt.Sprintf(
			"Requested capacity in Gb: %v for testrun (id: %v). Could not create file system bigger than %vGb\n",
			neededFsCapacityGb,
			req.TestRunId,
			MAX_FILESTORE_CAPACITY_GB,
		))
	}

	var fsInstanceConfig *file.Instance = h.buildFSInstanceConfig(neededFsCapacityGb)
	createFsInstanceCall := h.CloudFileStoreService.Projects.Locations.Instances.Create(
		fmt.Sprintf("projects/%s/locations/%s", config.GetConfig().GoogleProjectID, config.GetConfig().ResourcesDeployLocations),
		fsInstanceConfig,
	)

	var instanceIdMd5 [16]byte = md5.Sum([]byte(fmt.Sprintf("%v-%v-%v", time.Now().Unix(), neededFsCapacityGb, req.TestRunId)))
	createFsInstanceCall.InstanceId(fmt.Sprintf("instance-%x", instanceIdMd5))

	createOperation, err := createFsInstanceCall.Do()
	if err != nil {
		return err
	}

	go h.executePostFSCreateOperationSteps(req.TestRunId, createOperation)

	return nil
}

func (h *Handler) executePostFSCreateOperationSteps(testRunId uint32, op *file.Operation) {
	finishedOperation, err := h.waitForFSOperationToFinish(op)
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

	// send data to run controller service
	fsCreatedEventData, _ := json.Marshal(&runControllerProto.FileSystemCreateEventData{TestRunId: testRunId})
	h.SendRunStateEvent(context.Background(), "filesystem.created", fsCreatedEventData)
}

func (h *Handler) waitForFSOperationToFinish(op *file.Operation) (*file.Operation, error) {
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
			return nil, errors.New(fmt.Sprintf("FS operation %s failed with error: %v", op.Name, operation.Error))
		}

		return operation, nil
	}
}

func (h *Handler) buildFSInstanceConfig(capacityGb int64) *file.Instance {
	return &file.Instance{
		Description: "",
		FileShares: []*file.FileShareConfig{
			&file.FileShareConfig{
				CapacityGb: capacityGb,
				Name:       "target",
			},
		},
		Tier: "BASIC_HDD",
		Networks: []*file.NetworkConfig{
			&file.NetworkConfig{
				Modes:           []string{"MODE_IPV4"},
				Network:         "default",
				ReservedIpRange: "",
			},
		},
	}
}
