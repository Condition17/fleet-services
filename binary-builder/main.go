package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/Condition17/fleet-services/binary-builder/proto/binary-builder"
)

type binaryBuilderServer struct {
	proto.UnimplementedBinaryBuilderServer
}

func (s *binaryBuilderServer) Hello(ctx context.Context, req *proto.EmptyMessage) (*proto.EmptyMessage, error) {
	fmt.Println("Executed hello! <3")
	return &proto.EmptyMessage{}, nil
}

const CHUNK_HANDLERS_COUNT = 10
const CHUNKS_QUEUE_SIZE = 100

type ChunkData struct {
	fileOffset uint
}

var done chan bool = make(chan bool)

var chunksQueue chan ChunkData = make(chan ChunkData, CHUNKS_QUEUE_SIZE)

func handleChunkData(handlerIndex int) {
	for chunkDetails := range chunksQueue {
		time.Sleep(2 * time.Second)
		fmt.Printf("[%d] Handled chunk details: %v\n", handlerIndex, chunkDetails.fileOffset)
	}
}

func main() {

	for i := 0; i < CHUNK_HANDLERS_COUNT; i++ {
		go handleChunkData(i)
	}
	// Simulate a request
	fmt.Println("Mount volume")
	fmt.Println("Create target file and get descriptor")
	for i := 0; i <= 800; i++ {
		chunksQueue <- ChunkData{fileOffset: uint(i)}
		fmt.Println("Added chunk offset:", i)
	}
	<-done

	// simulate request start
	var testRunId uint32 = 254

}
