package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	topics "github.com/Condition17/fleet-services/lib/communication"
	resourceManagerProto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	nfsModule "github.com/Condition17/fleet-services/river-runner/nfs"
	proto "github.com/Condition17/fleet-services/river-runner/proto/river-runner"
	riverSdk "github.com/Condition17/fleet-services/river/sdk"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerPb "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"google.golang.org/grpc"
	"log"
	"path"
)

type Handler struct {
	proto.UnimplementedRiverRunnerServer
	resourceManagerClient resourceManagerProto.ResourceManagerServiceClient
	fileServiceClient     fileServiceProto.FileServiceClient
	runStateTopic         *pubsub.Topic
}

func NewHandler(externalServicesConn *grpc.ClientConn, pubSubClient *pubsub.Client) *Handler {
	return &Handler{
		resourceManagerClient: resourceManagerProto.NewResourceManagerServiceClient(externalServicesConn),
		fileServiceClient:     fileServiceProto.NewFileServiceClient(externalServicesConn),
		runStateTopic:         pubSubClient.Topic(topics.RunStateTopic),
	}
}

func (h *Handler) RunRiver(ctx context.Context, req *proto.RunRequest) (*proto.EmptyResponse, error) {
	var fileSystemData *resourceManagerProto.FileSystem
	var fileData *fileServiceProto.File
	// get file system details for test run
	if res, err := h.resourceManagerClient.GetFileSystem(ctx, &resourceManagerProto.FileSystemSpec{TestRunId: req.TestRunId}); err != nil {
		log.Printf("Error encountered while retrieving file system details for the provided test run (id: %v): %v\n", req.TestRunId, err)
		return nil, errors.New(fmt.Sprintf("Error retrieving file system details for test run (id: %v): %v\n", req.TestRunId, err))
	} else {
		fileSystemData = res.FileSystem
	}
	log.Println("Retrieved file system details:", fileSystemData)

	// get file details
	if res, err := h.fileServiceClient.ReadFile(context.Background(), &fileServiceProto.File{Id: fileSystemData.TestRun.FileId}); err != nil {
		log.Printf("Error  retrieving file details for provided test run (id: %v): %v\n", req.TestRunId, err)
		return nil, errors.New(fmt.Sprintf("Error retrieving file data for test run (id: %v): %v\n", req.TestRunId, err))
	} else {
		fileData = res.File
	}
	log.Println("Retrieved file details:", fileData)

	// mount volume
	mountDirPath := path.Join("/mnt/", fmt.Sprintf("testrun_%v", req.TestRunId))
	if err := nfsModule.MountVolume(fileSystemData.IP, fileSystemData.FileShareName, mountDirPath); err != nil {
		log.Printf("Error mounting volume: %v", err)
		return nil, errors.New(fmt.Sprintf("Error mounting volume: %v", err))
	}

	// async start running River on given file
	// this block also sends run results to run state queue
	go func() {
		out, err := riverSdk.Run(path.Join(mountDirPath, fileData.Name),
			"-secondsBetweenStats", "2",
			"-arch", "x64",
			"-max", "1",
			"-outputType", "textual",
		)

		if err != nil {
			// TODO: send async error message
			log.Printf("River run command error: %v", err)
			_ = nfsModule.UmountVolume(mountDirPath)
			return
		}
		log.Printf("River run successfully finished with output: %s\n", out)

		// unmount volume
		if err = nfsModule.UmountVolume(mountDirPath); err != nil {
			// TODO: send async error message
			return
		}
		// construct and send the notification message
		eventData, _ := json.Marshal(&runControllerPb.RiverRunFinishedEventData{TestRunId: req.TestRunId})
		_ = h.sendRunStateEvent(context.Background(), runStateEvents.TEST_RUN_FINISHED, eventData)
	}()

	return &proto.EmptyResponse{}, nil
}

func (h *Handler) sendRunStateEvent(ctx context.Context, eventType string, data []byte) error {
	msg, _ := json.Marshal(&runControllerPb.Event{
		Type: eventType,
		Meta: &runControllerPb.EventMetadata{Authorization: []byte("")},
		Data: data,
	})
	result := h.runStateTopic.Publish(ctx, &pubsub.Message{Data: msg})
	id, err := result.Get(ctx)

	if err != nil {
		log.Println("Error encountered sending message to run controller service:", err)
		return err
	}
	log.Printf("Published message to '%v' topic. Message ID: %v\n", h.runStateTopic.String(), id)

	return nil
}
