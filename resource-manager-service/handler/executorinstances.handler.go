package handler

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"google.golang.org/api/compute/v1"

	"github.com/Condition17/fleet-services/resource-manager-service/config"
	"github.com/Condition17/fleet-services/resource-manager-service/model"
	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
)

var (
	CONFIGS          config.Config = config.GetConfig()
	StartupScriptUrl string        = "https://storage.googleapis.com/fleet-metadata/setup.sh"
	MachineType      string        = fmt.Sprintf("zones/%s/machineTypes/e2-small", CONFIGS.ResourcesDeployLocations)
	DiskSizeGb       int64         = 10
	DiskSourceImage  string        = "projects/debian-cloud/global/images/debian-10-buster-v20201112"
)

func (h *Handler) ProvisionExecutorInstance(ctx context.Context, req *proto.ExecutorInstanceSpec, res *proto.EmptyResponse) error {
	var computeInstanceConfig *compute.Instance = h.buildComputeInstanceConfig(fmt.Sprintf("%v-%v", time.Now().Unix(), req.TestRunId))
	insertOperation, err := h.CloudComputeEngineService.Instances.Insert(CONFIGS.GoogleProjectID, CONFIGS.ResourcesDeployLocations, computeInstanceConfig).Do()
	if err != nil {
		return err
	}

	go h.executePostInstanceInsertOperationSteps(req.TestRunId, insertOperation)

	return nil
}

func (h *Handler) executePostInstanceInsertOperationSteps(testRunId uint32, op *compute.Operation) {
	finishedOperation, err := h.waitForComputeOperationToFinish(op)
	if err != nil {
		fmt.Printf("Error encountered while waiting compute instance insert operation finish: %v", err)
		h.SendServiceError(context.Background(), testRunId, err)
		return
	}

	createdInstance, err := h.CloudComputeEngineService.Instances.Get(CONFIGS.GoogleProjectID, CONFIGS.ResourcesDeployLocations, fmt.Sprintf("%v", finishedOperation.TargetId)).Do()
	if err != nil {
		fmt.Printf("Error encountered while retrieving created instance details: %v", err)
		h.SendServiceError(context.Background(), testRunId, err)
		return
	}

	_, err = h.ExecutorInstanceRepository.Create(&model.ExecutorInstance{
		IP:          createdInstance.NetworkInterfaces[0].NetworkIP,
		Name:        createdInstance.Name,
		MachineType: createdInstance.MachineType,
		TestRunID:   testRunId,
	})

	if err != nil {
		fmt.Printf("Error encountered on executor instance db entity creation: %v", err)
		h.SendServiceError(context.Background(), testRunId, err)
		return
	}

	// send data to run controller service
	execInstanceCreatedEventData, _ := json.Marshal(&runControllerProto.ExecutorInstanceProvisionedEventData{TestRunId: testRunId})
	h.SendRunStateEvent(context.Background(), runStateEvents.ExecutorInstanceProvisioned, execInstanceCreatedEventData)
}

func (h *Handler) waitForComputeOperationToFinish(op *compute.Operation) (*compute.Operation, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		operationStatusCall := h.CloudComputeEngineService.ZoneOperations.Get(
			CONFIGS.GoogleProjectID,
			CONFIGS.ResourcesDeployLocations,
			op.Name,
		)
		operation, err := operationStatusCall.Do()
		if err != nil {
			return nil, err
		}

		if operation.Status != "DONE" {
			continue
		}

		// the operation has finished
		ticker.Stop()
		if operation.Error != nil {
			return nil, errors.New(fmt.Sprintf("Compute operation %s failed with error: %v", op.Name, operation.Error))
		}

		return operation, nil
	}
}

func (h *Handler) buildComputeInstanceConfig(name string) *compute.Instance {
	var instanceNameMd5Seq [16]byte = md5.Sum([]byte(fmt.Sprintf("%v-%v", time.Now().Unix(), name)))

	return &compute.Instance{
		Description: "",
		Name:        fmt.Sprintf("instance-%x", instanceNameMd5Seq),
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				&compute.MetadataItems{
					Key:   "startup-script-url",
					Value: &StartupScriptUrl,
				},
			},
		},
		MachineType: MachineType,
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				AutoDelete: true,
				Boot:       true,
				Mode:       "READ_WRITE",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskSizeGb:  DiskSizeGb,
					SourceImage: DiskSourceImage,
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{},
		},
	}
}
