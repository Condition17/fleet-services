package handler

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"math"
	"time"

	"github.com/Condition17/fleet-services/resource-manager-service/config"
	"github.com/Condition17/fleet-services/resource-manager-service/model"
	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	microErrors "github.com/micro/go-micro/v2/errors"
	"google.golang.org/api/file/v1"
	"gorm.io/gorm"
)

const MinFilestoreCapacityGb int64 = 1024
const MaxFilestoreCapacityGb int64 = 63900

func (h *Handler) GetFileSystem(ctx context.Context, req *proto.FileSystemSpec, res *proto.FileSystemDetails) error {
	result, err := h.FileSystemRepository.GetByTestRunId(req.TestRunId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), fmt.Sprintf("%v", err))
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.FileSystem = model.UnmarshalFileSystem(result)

	return nil
}

func (h *Handler) ProvisionFileSystem(ctx context.Context, req *proto.FileSystemSpec, res *proto.EmptyResponse) error {
	var requestedSizeGb int64 = int64(math.Round(float64(req.SizeInBytes)/float64(math.Pow10(9)) + 0.5))
	var neededFsCapacityGb = int64(math.Max(float64(MinFilestoreCapacityGb), float64(requestedSizeGb)))
	var locationPath = fmt.Sprintf("projects/%s/locations/%s", config.GetConfig().GoogleProjectID, config.GetConfig().ResourcesDeployLocations)

	if neededFsCapacityGb > MaxFilestoreCapacityGb {
		return errors.New(fmt.Sprintf(
			"Requested capacity in Gb: %v for testrun (id: %v). Could not create file system bigger than %vGb\n",
			neededFsCapacityGb,
			req.TestRunId,
			MaxFilestoreCapacityGb,
		))
	}

	alreadyCreatedInstance, err := h.getCreatedFSInstance(locationPath)
	if alreadyCreatedInstance != nil {
		go h.createFsEntryInDb(req.TestRunId, alreadyCreatedInstance)
		return nil
	}

	var fsInstanceConfig *file.Instance = h.buildFSInstanceConfig(neededFsCapacityGb)
	createFsInstanceCall := h.CloudFileStoreService.Projects.Locations.Instances.Create(locationPath, fsInstanceConfig)

	var instanceIdMd5 [16]byte = md5.Sum([]byte(fmt.Sprintf("%v-%v-%v", time.Now().Unix(), neededFsCapacityGb, req.TestRunId)))
	createFsInstanceCall.InstanceId(fmt.Sprintf("instance-%x", instanceIdMd5))

	createOperation, err := createFsInstanceCall.Do()
	if err != nil {
		return err
	}

	go h.executePostFSCreateOperationSteps(req.TestRunId, createOperation)

	return nil
}

func (h *Handler) getCreatedFSInstance(locationPath string) (*file.Instance, error){
	fsInstancesListCall := h.CloudFileStoreService.Projects.Locations.Instances.List(locationPath)
	fsInstanceListResponse, err := fsInstancesListCall.Do()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while listing filestore instances: %v", err))
	}

	if len(fsInstanceListResponse.Instances) < 2 {
		return nil, nil
	}

	for _, instance := range fsInstanceListResponse.Instances {
		if instance.State == "READY" {
			return instance, nil
		}
	}

	return nil, nil
}

func (h *Handler) executePostFSCreateOperationSteps(testRunId uint32, op *file.Operation) {
	finishedOperation, err := h.waitForFSOperationToFinish(op)
	if err != nil {
		fmt.Printf("Error encountered while waiting filestore create operation finish %v", err)
		h.SendServiceError(context.Background(), testRunId, err)
		return
	}

	var fsInstance *file.Instance
	if err := json.Unmarshal(finishedOperation.Response, &fsInstance); err != nil {
		fmt.Printf("Error encountered while unmarshalling operation response: %v", err)
		h.SendServiceError(context.Background(), testRunId, err)
		return
	}

	h.createFsEntryInDb(testRunId, fsInstance)
}

func (h *Handler) createFsEntryInDb(testRunId uint32, fsInstance *file.Instance) {
	_, err := h.FileSystemRepository.Create(&model.FileSystem{
		IP:                  fsInstance.Networks[0].IpAddresses[0],
		Name:                fsInstance.Name,
		FileShareCapacityGb: fsInstance.FileShares[0].CapacityGb,
		FileShareName:       fsInstance.FileShares[0].Name,
		TestRunID:           testRunId,
	})

	if err != nil {
		fmt.Printf("Error encountered on file system db entity creation: %v", err)
		h.SendServiceError(context.Background(), testRunId, err)
		return
	}

	// send data to run controller service
	fsProvisionedEventData, _ := json.Marshal(&runControllerProto.FileSystemProvisionedEventData{TestRunId: testRunId})
	h.SendRunStateEvent(context.Background(), runStateEvents.FileSystemProvisioned, fsProvisionedEventData)
	return
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
