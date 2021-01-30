package cloudstorage

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

const bucketName string = "fleet-files-chunks"

type GcsClient struct {
	storage *storage.Client
}

func InitClient() (*GcsClient, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &GcsClient{storage: client}, nil
}

func (gcsClient *GcsClient) UploadChunk(chunkName string, content string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()

	bucket := gcsClient.storage.Bucket(bucketName)
	obj := bucket.Object(chunkName)
	w := obj.NewWriter(ctx)

	if _, err := fmt.Fprintf(w, content); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		fmt.Println("Error encountered writing chunk :", err)
		return err
	}

	return nil
}
