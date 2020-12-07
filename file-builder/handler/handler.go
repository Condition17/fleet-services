package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	chunksStorage "github.com/Condition17/fleet-services/file-builder/chunks-storage"
	fileComposer "github.com/Condition17/fleet-services/file-builder/composer"
	nfsModule "github.com/Condition17/fleet-services/file-builder/nfs"
	proto "github.com/Condition17/fleet-services/file-builder/proto/file-builder"
	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	topics "github.com/Condition17/fleet-services/lib/communication"
	resourceManagerProto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"google.golang.org/grpc"
	"log"
	"path"
)

type Handler struct {
	proto.UnimplementedFileBuilderServer
	serviceName           string
	resourceManagerClient resourceManagerProto.ResourceManagerServiceClient
	fileServiceClient     fileServiceProto.FileServiceClient
	fileComposer          *fileComposer.Composer
	runStateTopic         *pubsub.Topic
}

func NewHandler(externalServicesConn *grpc.ClientConn, pubSubClient pubsub.Client, chunksStorageClient chunksStorage.Client) *Handler {
	fileServiceClient := fileServiceProto.NewFileServiceClient(externalServicesConn)

	return &Handler{
		serviceName:           "File builder",
		fileComposer:          fileComposer.NewComposer(chunksStorageClient, fileServiceClient),
		resourceManagerClient: resourceManagerProto.NewResourceManagerServiceClient(externalServicesConn),
		fileServiceClient:     fileServiceClient,
		runStateTopic:         pubSubClient.Topic(topics.RunStateTopic),
	}
}

func (h *Handler) AssembleFile(ctx context.Context, req *proto.FileAssembleRequest) (*proto.EmptyResponse, error) {
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

	// create target directory to mount files to
	mountDirPath := path.Join("/mnt/", fmt.Sprintf("testrun_%v", req.TestRunId))
	if err := nfsModule.MountVolume(fileSystemData.IP, fileSystemData.FileShareName, mountDirPath); err != nil {
		log.Printf("Error mounting volume: %v", err)
		return nil, errors.New(fmt.Sprintf("Error mounting volume: %v", err))
	}

	go func() {
		feedback := h.fileComposer.ComposeFile(&fileComposer.FileSpec{
			Id:               fileData.Id,
			Name:             fileData.Name,
			TotalChunksCount: fileData.TotalChunksCount,
			MaxChunkSize:     fileData.MaxChunkSize,
			ParentDir:        mountDirPath,
		})
		select {
		case <-feedback.SuccessChan:
			log.Println("File successfully assembled")
			// TODO: THIS IS FOR TESTING PURPOSES - REMOVE IT
			_ = h.sendServiceError(context.Background(), req.TestRunId, errors.New("file successfully assembled"))
		case err := <-feedback.ErrorChan:
			log.Printf("[SERVICE ERROR] Error encountered assembling file (id: %v): %v", fileData.Id, err)
			_ = h.sendServiceError(context.Background(), req.TestRunId, err)
			_ = nfsModule.UmountVolume(mountDirPath)
			return
		}

		// unmount volume
		if err := nfsModule.UmountVolume(mountDirPath); err != nil {
			log.Printf("[SERVICE ERROR] Error encountered umounting volume (path: %v): %v", mountDirPath, err)
			_ = h.sendServiceError(context.Background(), req.TestRunId, err)
			return
		}
		// construct and send the notification message
		eventData, _ := json.Marshal(&runControllerProto.FileAssemblySucceededEventData{TestRunId: req.TestRunId})
		_ = h.sendRunStateEvent(context.Background(), runStateEvents.FileAssemblySuccess, eventData)
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
		log.Println("Error sending message to run controller service:", err)
		return err
	}
	log.Printf("Published message to '%v' topic. Message ID: %v\n", h.runStateTopic.String(), id)

	return nil
}
