package main

import (
	"context"
	"fmt"
	"github.com/Condition17/fleet-services/binary-builder/config"
	"google.golang.org/grpc"
	"log"
	"time"

	proto "github.com/Condition17/fleet-services/binary-builder/proto/binary-builder"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	resourceManagerPb "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
)

type binaryBuilderServer struct {
	proto.UnimplementedBinaryBuilderServer
}

func (s *binaryBuilderServer) Hello(ctx context.Context, req *proto.EmptyMessage) (*proto.EmptyMessage, error) {
	fmt.Println("Executed hello! <3")
	return &proto.EmptyMessage{}, nil
}

const ChunkHandlersCount = 10
const ChunksQueueSize = 100

type ChunkData struct {
	fileOffset uint
}

var done chan bool = make(chan bool)

var chunksQueue chan ChunkData = make(chan ChunkData, ChunksQueueSize)

func handleChunkData(handlerIndex int) {
	for chunkDetails := range chunksQueue {
		time.Sleep(2 * time.Second)
		fmt.Printf("[%d] Handled chunk details: %v\n", handlerIndex, chunkDetails.fileOffset)
	}
}

func main() {
	//for i := 0; i < ChunkHandlersCount; i++ {
	//	go handleChunkData(i)
	//}
	//// Simulate a request
	//fmt.Println("Mount volume")
	//fmt.Println("Create target file and get descriptor")
	//for i := 0; i <= 800; i++ {
	//	chunksQueue <- ChunkData{fileOffset: uint(i)}
	//	fmt.Println("Added chunk offset:", i)
	//}
	//<-done

	// simulate request start

	configs := config.GetConfig()
	conn, err := grpc.Dial(configs.FleetServicesGrpcProxyUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error encountered while creating connection to fleet services grpc proxy: %v", err)
	}
	log.Println("Connection to fleet services GRPC proxy initiated");
	defer conn.Close()

	var fileServiceClient fileServicePb.FileServiceClient = fileServicePb.NewFileServiceClient(conn)
	var resourceManagerClient resourceManagerPb.ResourceManagerServiceClient = resourceManagerPb.NewResourceManagerServiceClient(conn)

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

	for chunkIndex:= uint64(0); chunkIndex < fileDetails.TotalChunksCount; chunkIndex++ {
		res, err := fileServiceClient.GetChunkDetailsByIndexInFile(context.Background(), &fileServicePb.ChunkSpec{FileId: fileDetails.Id, Index: chunkIndex})
		if err != nil {
			log.Fatalf("Error encountered while retrieving chunk (index: %v) details for file (id: %v): %v", chunkIndex, fileDetails.Id, err)
		}
		fmt.Println("Chunk details:", res)
	}
}
