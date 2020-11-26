package main

import (
	"context"
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
)

//type binaryBuilderServer struct {
//	proto.UnimplementedBinaryBuilderServer
//}
//
//func (s *binaryBuilderServer) Hello(ctx context.Context, req *proto.EmptyMessage) (*proto.EmptyMessage, error) {
//	fmt.Println("Executed hello! <3")
//	return &proto.EmptyMessage{}, nil
//}

const ChunkHandlersCount = 10
const ChunksQueueSize = 100
const ChunksBucketName = "fleet-files-chunks"
var mu sync.Mutex = sync.Mutex{}

type ChunkDetails struct {
	Sha2 string
	File *os.File
	FileOffset uint64
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
		_, err = chunkDetails.File.Seek(int64(chunkDetails.FileOffset), 0)
		if err != nil {
			log.Printf("Error while seeking file offset (offset: %v): %v\n", chunkDetails.FileOffset, err)
			mu.Unlock()
			continue
		}
		if _, err := io.Copy(chunkDetails.File, chunkBytesReader); err != nil {
			log.Printf("Error while dowloading chunk (key: %v) bytes: %v\n", chunkDetails.Sha2, err)
			mu.Unlock()
			continue
		}
		log.Printf("Successfully written chunk '%v' at file offset '%v'\n", chunkDetails.Sha2, chunkDetails.FileOffset)
		mu.Unlock()
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


	for i := uint64(0); i < fileDetails.TotalChunksCount; i++ {
		res, err := fileServiceClient.GetChunkDetailsByIndexInFile(context.Background(), &fileServicePb.ChunkSpec{FileId: fileDetails.Id, Index: i})
		if err != nil {
			log.Fatalf("Error encountered while retrieving chunk (index: %v) details for file (id: %v): %v", i, fileDetails.Id, err)
		}
		chunksQueue <- ChunkDetails{Sha2: res.Chunk.Sha2, File: file, FileOffset: i * uint64(fileDetails.MaxChunkSize)}
	}

	// TODO: remove this after finishing the server layout
	<- done
}
