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
	"github.com/Condition17/fleet-services/river-runner/config"
	nfsModule "github.com/Condition17/fleet-services/river-runner/nfs"
	proto "github.com/Condition17/fleet-services/river-runner/proto/river-runner"
	riverSdk "github.com/Condition17/fleet-services/river/sdk"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"google.golang.org/grpc"
	"log"
	"path"
)

type Handler struct {
	proto.UnimplementedRiverRunnerServer
	serviceName           string
	resourceManagerClient resourceManagerProto.ResourceManagerServiceClient
	fileServiceClient     fileServiceProto.FileServiceClient
	runStateTopic         *pubsub.Topic
}

func NewHandler(externalServicesConn *grpc.ClientConn, pubSubClient *pubsub.Client) *Handler {
	return &Handler{
		serviceName:           "RiverRunner",
		resourceManagerClient: resourceManagerProto.NewResourceManagerServiceClient(externalServicesConn),
		fileServiceClient:     fileServiceProto.NewFileServiceClient(externalServicesConn),
		runStateTopic:         pubSubClient.Topic(topics.RunStateTopic),
	}
}

func (h *Handler) RunRiver(ctx context.Context, req *proto.RunRequest) (*proto.EmptyResponse, error) {
	var fileSystemData *resourceManagerProto.FileSystem
	var fileData *fileServiceProto.File
	configs := config.GetConfig()
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
		log.Println("Evaluating file with River...")
		exitCode, err := riverSdk.Run(path.Join(mountDirPath, fmt.Sprintf("/%v", req.TestRunId), fileData.Name),
			"-secondsBetweenStats", "2",
			"-arch", "x64",
			"-max", "1000000",
			"-outputType", "textual",
			"-outputEndpoint",
			fmt.Sprintf("%v/testRunService/RegisterRunIssue?testRunId=%v", configs.FleetServicesHttpApiUrl, req.TestRunId),
		)

		// unmount volume
		if err = nfsModule.UmountVolume(mountDirPath); err != nil {
			log.Printf("[SERVICE ERROR] Error encountered umounting volume (path: %v): %v", mountDirPath, err)
			_ = h.sendServiceError(context.Background(), req.TestRunId, err)
			return
		}
		// construct and send the notification message
		eventData, _ := json.Marshal(&runControllerProto.FileEvaluationFinishedEventData{TestRunId: req.TestRunId, ExitCode: uint32(exitCode)})
		_ = h.sendRunStateEvent(context.Background(), runStateEvents.FileEvaluationFinished, eventData)
	}()

	return &proto.EmptyResponse{}, nil
}

func (h *Handler) sendServiceError(ctx context.Context, testRunId uint32, err error) error {
	eventData, _ := json.Marshal(&runControllerProto.ServiceErrorEventData{Source: h.serviceName, TestRunId: testRunId, Error: []byte(err.Error())})
	return h.sendRunStateEvent(ctx, runStateEvents.ServiceError, eventData)
}

func (h *Handler) sendRunStateEvent(ctx context.Context, eventType string, data []byte) error {
	msg, _ := json.Marshal(&runControllerProto.Event{
		Type: eventType,
		Meta: &runControllerProto.EventMetadata{Authorization: []byte("")},
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
