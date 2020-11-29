package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/Condition17/fleet-services/file-builder/chunks-storage"
	"github.com/Condition17/fleet-services/file-builder/config"
	proto "github.com/Condition17/fleet-services/file-builder/proto/file-builder"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	resourceManagerPb "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	runControllerPb "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
)

const runStateTopic = "test-run-state"
const (
	ChunkHandlersCount = 10
	ChunksQueueSize    = 100
	ChunksBucketName   = "fleet-files-chunks"
)

type FileBuildEventType string

const (
	chunkProcessed       FileBuildEventType = "chunkProcessed"
	chunkProcessingError FileBuildEventType = "chunkProcessingError"
)

var mu sync.Mutex = sync.Mutex{}

type ChunkDetails struct {
	BuildFile  BuildFile
	Sha2       string
	FileOffset uint64
}

type BuildFile struct {
	MountVolumePath  string
	File             *os.File
	TestRunId        uint32
	Spec             *fileServicePb.File
	BuildUpdatesChan chan FileBuildEvent
}

type FileBuildEvent struct {
	Type    FileBuildEventType
	Error   error
	Payload ChunkDetails
}

var chunksQueue chan ChunkDetails = make(chan ChunkDetails, ChunksQueueSize)

// GLOBALS -- todo: fix this
var (
	testRunStateTopic     *pubsub.Topic
	fileServiceClient     fileServicePb.FileServiceClient
	resourceManagerClient resourceManagerPb.ResourceManagerServiceClient
)

type fileBuilderServer struct {
	proto.UnimplementedFileBuilderServer
}

func (s *fileBuilderServer) TestCall(ctx context.Context, req *proto.FileAssembleRequest) (*proto.EmptyResponse, error) {
	mountPointAddr := "10.41.254.146"// TODO: fix this
	mountPointSource := ":/target"
	mountDirPath := path.Join("/mnt/", fmt.Sprintf("testrun_%v", req.TestRunId))
	if err := syscall.Mount(mountPointSource, mountDirPath, "nfs", 0, fmt.Sprintf("nolock,addr=%s", mountPointAddr)); err != nil {
		log.Fatalf("Syscall mount error: %v", err)
	}
	fmt.Println("NFS successfully mounted.")

	// try file creation in NFS
	f, err := os.OpenFile(filepath.Join(mountDirPath, "program_file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	fmt.Println("Trying to unmount")
	out, err := exec.Command("umount", "-l", mountDirPath).Output()
	if err != nil {
		log.Fatalf("Error unmounting fs: %s | Out: %s", err, out)
	}
	fmt.Println("Successfully unmounted")
	return &proto.EmptyResponse{}, nil
}

func (s *fileBuilderServer) AssembleFile(ctx context.Context, req *proto.FileAssembleRequest) (*proto.EmptyResponse, error) {
	testRunId := req.TestRunId
	// get file system details
	response, err := resourceManagerClient.GetFileSystem(context.Background(), &resourceManagerPb.FileSystemSpec{TestRunId: testRunId})
	if err != nil {
		// TODO: return data here to caller
		log.Fatalf("Error encountered while retrieving file system details for the provided test run (id: %v): %v", testRunId, err)
	}
	var fileSystemDetails *resourceManagerPb.FileSystem = response.FileSystem

	// get file details
	res, err := fileServiceClient.ReadFile(context.Background(), &fileServicePb.File{Id: fileSystemDetails.TestRun.FileId})
	if err != nil {
		// TODO: return data to the caller here
		log.Fatalf("Error encountered while retrieving file details for provided test run (id: %v): %v", testRunId, err)
	}
	var fileDetails *fileServicePb.File = res.File
	// the associated file was found

	// --- mount volume in mount directory
	var fileDir string = filepath.Join("./", "mnt/target")

	// -- create target file - this will be built assembling the downloaded chunks
	file, _ := os.OpenFile(filepath.Join(fileDir, fileDetails.Name), os.O_CREATE|os.O_RDWR, 0666)
	var buildFile BuildFile = BuildFile{File: file, MountVolumePath: fileDir, TestRunId: testRunId, Spec: fileDetails, BuildUpdatesChan: make(chan FileBuildEvent)}
	go handleExecutionFeedback(buildFile, testRunStateTopic)
	for i := uint64(0); i < fileDetails.TotalChunksCount; i++ {
		res, err := fileServiceClient.GetChunkDetailsByIndexInFile(context.Background(), &fileServicePb.ChunkSpec{FileId: fileDetails.Id, Index: i})
		if err != nil {
			log.Fatalf("Error encountered while retrieving chunk (index: %v) details for file (id: %v): %v", i, fileDetails.Id, err)
		}
		chunksQueue <- ChunkDetails{BuildFile: buildFile, Sha2: res.Chunk.Sha2, FileOffset: i * uint64(fileDetails.MaxChunkSize)}
	}

	return &proto.EmptyResponse{}, nil
}

func handleChunkDetails(chunksStorageClient *chunksStorage.Client) {
	for chunkDetails := range chunksQueue {
		chunkBytesReader, err := chunksStorageClient.GetObjectReader(chunkDetails.Sha2)
		if err != nil {
			log.Printf("Error encountered obtaining object reader for chunk (key: %v): %v\n", chunkDetails.Sha2, err)
			continue
		}
		// lock
		mu.Lock()
		log.Printf("Writing chunk '%v' at file offset '%v'\n", chunkDetails.Sha2, chunkDetails.FileOffset)
		_, err = chunkDetails.BuildFile.File.Seek(int64(chunkDetails.FileOffset), 0)
		if err != nil {
			chunkDetails.BuildFile.BuildUpdatesChan <- FileBuildEvent{Type: chunkProcessingError, Error: err, Payload: chunkDetails}
			mu.Unlock()
			continue
		}
		if _, err := io.Copy(chunkDetails.BuildFile.File, chunkBytesReader); err != nil {
			chunkDetails.BuildFile.BuildUpdatesChan <- FileBuildEvent{Type: chunkProcessingError, Error: err, Payload: chunkDetails}
			mu.Unlock()
			continue
		}
		log.Printf("Successfully written chunk '%v' at file offset '%v'\n", chunkDetails.Sha2, chunkDetails.FileOffset)
		mu.Unlock()
		chunkDetails.BuildFile.BuildUpdatesChan <- FileBuildEvent{Type: chunkProcessed, Payload: chunkDetails}
	}
}

func handleExecutionFeedback(buildFile BuildFile, runStateTopic *pubsub.Topic) {
	var appendedChunksCount uint64 = 0
	for fileBuildEvent := range buildFile.BuildUpdatesChan {
		switch fileBuildEvent.Type {
		case chunkProcessed:
			appendedChunksCount++
			if appendedChunksCount == buildFile.Spec.TotalChunksCount {
				log.Printf("File %v completely assembled\n", buildFile.Spec.Name)
				// close file
				buildFile.File.Close()
				// unmount volume
				// TODO: implement this in linux environment
				// Notify run-state service

				// construct the notification message first

				eventData, _ := json.Marshal(
					&runControllerPb.FileAssemblySucceededEventData{
						TestRunId: buildFile.TestRunId,
					},
				)
				msg, _ := json.Marshal(
					&runControllerPb.Event{
						Type: "file.assemblySucceeded",
						Meta: &runControllerPb.EventMetadata{
							Authorization: []byte(""),
						},
						Data: eventData,
					},
				)
				// send notification
				result := runStateTopic.Publish(context.Background(), &pubsub.Message{Data: msg})
				id, err := result.Get(context.Background())
				if err != nil {
					log.Println("Error encountered sending message to run controller service:", err)
				}
				log.Printf("Published message to '%v' topic. Message ID: %v\n", runStateTopic.String(), id)
			}
		case chunkProcessingError:
			// TODO: some retry rule needed - also a retry mechanism should be implemented
			log.Printf("Error encountered while processing chunk (key: %v): %v\n", fileBuildEvent.Payload.Sha2, fileBuildEvent.Error)
			chunksQueue <- fileBuildEvent.Payload
			log.Printf("Chunk (key: %v) re-added to processing queue", fileBuildEvent.Payload.Sha2)
		}
	}
}

func main() {
	configs := config.GetConfig()

	// Server startup
	var serverAddr string = fmt.Sprintf("0.0.0.0:%v", configs.ServerPort)
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Server failed to listen on port ':%v'. Error encountered: %v\n", configs.ServerPort, err)
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	proto.RegisterFileBuilderServer(grpcServer, &fileBuilderServer{})

	conn, err := grpc.Dial(configs.FleetServicesGrpcProxyUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error encountered while creating connection to fleet services grpc proxy: %v", err)
	}
	log.Println("Connection to fleet services GRPC proxy initiated")
	defer conn.Close()

	// Create chunks storage client
	chunksStorageClient, err := chunksStorage.NewChunksStorageClient(configs.GoogleProjectID, ChunksBucketName)
	if err != nil {
		log.Fatalf("Error encountered while initializing the chunks storage client: %v", err)
	}

	// Create pubsub client
	pubSubClient, err := pubsub.NewClient(context.Background(), configs.GoogleProjectID)
	if err != nil {
		log.Fatalf("Error on pubsub.NewClient: %v", err)
	}
	// Get test run state topic
	testRunStateTopic = pubSubClient.Topic(runStateTopic)
	fileServiceClient = fileServicePb.NewFileServiceClient(conn)
	resourceManagerClient = resourceManagerPb.NewResourceManagerServiceClient(conn)

	// Simulate server setup
	for i := 0; i < ChunkHandlersCount; i++ {
		go handleChunkDetails(chunksStorageClient)
	}

	fmt.Printf("Starting GRPC server on localhost:%v\n", configs.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
