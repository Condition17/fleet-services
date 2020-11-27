package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"io"
	"sync"

	//"fmt"
	"github.com/Condition17/fleet-services/binary-builder/config"
	"google.golang.org/grpc"
	//"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Condition17/fleet-services/binary-builder/chunks-storage"
	//proto "github.com/Condition17/fleet-services/binary-builder/proto/binary-builder"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	resourceManagerPb "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	runControllerPb "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
)

//type binaryBuilderServer struct {
//	proto.UnimplementedBinaryBuilderServer
//}
//
//func (s *binaryBuilderServer) Hello(ctx context.Context, req *proto.EmptyMessage) (*proto.EmptyMessage, error) {
//	fmt.Println("Executed hello! <3")
//	return &proto.EmptyMessage{}, nil
//}

const runStateTopic = "test-run-state"
const (
	ChunkHandlersCount = 10
	ChunksQueueSize = 100
	ChunksBucketName = "fleet-files-chunks"
)

type FileBuildEventType string
const (
	chunkProcessed FileBuildEventType = "chunkProcessed"
	chunkProcessingError FileBuildEventType = "chunkProcessingError"
)

var mu sync.Mutex = sync.Mutex{}

type ChunkDetails struct {
	BuildFile BuildFile
	Sha2 string
	FileOffset uint64
}

type BuildFile struct {
	MountVolumePath string
	File *os.File
	TestRunId uint32
	Spec *fileServicePb.File
	BuildUpdatesChan chan FileBuildEvent
}

type FileBuildEvent struct {
	Type FileBuildEventType
	Error error
	Payload ChunkDetails
}

type PubSubMessage struct {
	Body interface{}
}

var done chan bool = make(chan bool)
var chunksQueue chan ChunkDetails = make(chan ChunkDetails, ChunksQueueSize)

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
	// Simulate server setup

	configs := config.GetConfig()

	conn, err := grpc.Dial(configs.FleetServicesGrpcProxyUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error encountered while creating connection to fleet services grpc proxy: %v", err)
	}
	log.Println("Connection to fleet services GRPC proxy initiated");
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
	testRunStateTopic := pubSubClient.Topic(runStateTopic)

	var fileServiceClient fileServicePb.FileServiceClient = fileServicePb.NewFileServiceClient(conn)
	var resourceManagerClient resourceManagerPb.ResourceManagerServiceClient = resourceManagerPb.NewResourceManagerServiceClient(conn)
	// Simulate server setup
	for i := 0; i < ChunkHandlersCount; i++ {
		go handleChunkDetails(chunksStorageClient)
	}

	// Simulate request start
	// --- handler shit here
	var testRunId uint32 = 254

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

	// TODO: remove this after finishing the server layout
	<- done
}
