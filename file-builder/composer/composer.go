package composer

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	chunksStorage "github.com/Condition17/fleet-services/file-builder/chunks-storage"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	"io"
	"log"
	"os"
	"path"
	"sync"
)

const (
	reProcessCountLimit   = 5
	chunksDownloaderCount = 10
	chunksQueueSize       = 100
)

type FileSpec struct {
	mu               *sync.Mutex
	ParentDir        string
	FileOnDisk       *os.File
	Id               string
	Name             string
	TotalChunksCount uint64
	MaxChunkSize     uint32
}

type ChunkDetails struct {
	ReProcessCount int
	Sha2           string
	Offset         uint64
	File           *FileSpec
}

type FileComposeEventType string

const (
	chunkProcessingSuccess FileComposeEventType = "chunkProcessingSuccess"
	chunkProcessingError   FileComposeEventType = "chunkProcessingError"
)

type FileComposeEvent struct {
	Type    FileComposeEventType
	Error   error
	Payload ChunkDetails
}

type ComposeOperationFeedback struct {
	SuccessChan chan struct{}
	ErrorChan   chan error
}

type Composer struct {
	chunksStorage       chunksStorage.Client
	fileServiceClient   fileServicePb.FileServiceClient
	chunksDetailsQueue  chan ChunkDetails
	operationUpdateChan map[*FileSpec]chan FileComposeEvent
	operationFeedback   map[*FileSpec]ComposeOperationFeedback
}

func NewComposer(chunksStorageClient chunksStorage.Client, fileServiceClient fileServicePb.FileServiceClient) *Composer {
	var composer *Composer = &Composer{
		chunksStorage:       chunksStorageClient,
		fileServiceClient:   fileServiceClient,
		chunksDetailsQueue:  make(chan ChunkDetails, chunksQueueSize),
		operationUpdateChan: make(map[*FileSpec]chan FileComposeEvent),
		operationFeedback:   make(map[*FileSpec]ComposeOperationFeedback),
	}

	// Start chunks downloader goroutines
	for i := 0; i < chunksDownloaderCount; i++ {
		go composer.runChunkDownloader()
	}

	return composer
}

func (c *Composer) ComposeFile(fileSpec *FileSpec) ComposeOperationFeedback {
	// Ensure requested target file is created before composing it
	fileSpec.FileOnDisk, _ = os.OpenFile(path.Join(fileSpec.ParentDir, fileSpec.Name), os.O_CREATE|os.O_RDWR, 0666)
	// Create channels needed in the compose processes
	c.operationUpdateChan[fileSpec] = make(chan FileComposeEvent)
	c.operationFeedback[fileSpec] = ComposeOperationFeedback{
		SuccessChan: make(chan struct{}),
		ErrorChan:   make(chan error),
	}

	go c.handleOperationUpdates(fileSpec)
	go c.runChunkDetailsProducer(fileSpec)

	return c.operationFeedback[fileSpec]
}

func (c *Composer) runChunkDetailsProducer(fileSpec *FileSpec) {
	// Get each file chunk details from DB and add them on the chunkDetails queue
	for i := uint64(0); i < fileSpec.TotalChunksCount; i++ {
		res, err := c.fileServiceClient.GetChunkDetailsByIndexInFile(context.Background(), &fileServicePb.ChunkSpec{FileId: fileSpec.Id, Index: i})
		if err != nil {
			c.operationUpdateChan[fileSpec] <- FileComposeEvent{
				Type:  chunkProcessingError,
				Error: errors.New(fmt.Sprintf("Error encountered while retrieving chunk (index: %v) details for file (id: %v): %v", i, fileSpec.Id, err)),
			}
			return
		}
		c.chunksDetailsQueue <- ChunkDetails{File: fileSpec, Sha2: res.Chunk.Sha2, Offset: i * uint64(fileSpec.MaxChunkSize)}
	}
}

func (c *Composer) runChunkDownloader() {
	var chunkBytesReader *storage.Reader
	chunkDetails := <-c.chunksDetailsQueue

	if !c.fileAvailableForComposing(chunkDetails.File) {
		c.runChunkDownloader()
	}

	chunkBytesReader, err := c.chunksStorage.GetObjectReader(chunkDetails.Sha2)
	if err != nil {
		c.operationUpdateChan[chunkDetails.File] <- FileComposeEvent{Type: chunkProcessingError, Error: err, Payload: chunkDetails}
		c.runChunkDownloader()
	}

	// get lock on chunk's associated file
	chunkDetails.File.mu.Lock()
	// re-check if file is still available for assembling process
	if !c.fileAvailableForComposing(chunkDetails.File) {
		c.runChunkDownloader()
	}

	log.Printf("Writing chunk '%v' at file offset '%v'\n", chunkDetails.Sha2, chunkDetails.Offset)
	if _, err := chunkDetails.File.FileOnDisk.Seek(int64(chunkDetails.Offset), 0); err != nil {
		c.operationUpdateChan[chunkDetails.File] <- FileComposeEvent{
			Type:    chunkProcessingError,
			Error:   errors.New(fmt.Sprintf("Error encountered obtaining object reader for chunk (key: %v): %v\n", chunkDetails.Sha2, err)),
			Payload: chunkDetails,
		}
		chunkDetails.File.mu.Unlock()
		c.runChunkDownloader()
	}
	if _, err := io.Copy(chunkDetails.File.FileOnDisk, chunkBytesReader); err != nil {
		c.operationUpdateChan[chunkDetails.File] <- FileComposeEvent{Type: chunkProcessingError, Error: err, Payload: chunkDetails}
		chunkDetails.File.mu.Unlock()
		c.runChunkDownloader()
	}
	log.Printf("Successfully written chunk '%v' at file offset '%v'\n", chunkDetails.Sha2, chunkDetails.Offset)
	chunkDetails.File.mu.Unlock()
	c.operationUpdateChan[chunkDetails.File] <- FileComposeEvent{Type: chunkProcessingSuccess, Payload: chunkDetails}
}

func (c *Composer) handleOperationUpdates(fileSpec *FileSpec) {
	var appendedChunksCount uint64 = 0
	composeEvent := <-c.operationUpdateChan[fileSpec]
	switch composeEvent.Type {
	case chunkProcessingSuccess:
		appendedChunksCount++
		if appendedChunksCount == fileSpec.TotalChunksCount {
			_ = fileSpec.FileOnDisk.Close()
			c.operationFeedback[fileSpec].SuccessChan <- struct{}{}
			delete(c.operationUpdateChan, fileSpec)
			delete(c.operationFeedback, fileSpec)
			return
		}
	case chunkProcessingError:
		log.Printf("Error encountered while processing chunk. Event details: %v\n", composeEvent)
		if (composeEvent.Payload != ChunkDetails{} && composeEvent.Payload.ReProcessCount < reProcessCountLimit) {
			composeEvent.Payload.ReProcessCount++
			c.chunksDetailsQueue <- composeEvent.Payload
			log.Printf("Chunk (key: %v) re-added to processing queue", composeEvent.Payload.Sha2)
		} else {
			log.Printf("File compose process aborted for file (id: %v)\n", composeEvent.Payload.File.Id)
			c.operationFeedback[fileSpec].ErrorChan <- composeEvent.Error
			delete(c.operationUpdateChan, fileSpec)
			delete(c.operationFeedback, fileSpec)
			return
		}
	}
	c.handleOperationUpdates(fileSpec)
}

func (c *Composer) fileAvailableForComposing(file *FileSpec) bool {
	return c.operationUpdateChan[file] != nil
}
