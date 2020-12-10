package googleStorage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"time"
)

type GcsClient struct {
	storage *storage.Client
}

func GetClient() (*GcsClient, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &GcsClient{storage: client}, nil
}

// This method returns uploaded bytes blob public url
func (gcsClient *GcsClient) UploadBytes(bucketName string, targetName string, data []byte) (string, error) {
	const baseStorageUrl string = "https://storage.googleapis.com"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	dataWriter := gcsClient.storage.Bucket(bucketName).Object(targetName).NewWriter(ctx)
	dataWriter.ContentType = "application/octet-stream"

	if _, err := dataWriter.Write(data); err != nil {
		return "", err
	}

	if err := dataWriter.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/%v/%v", baseStorageUrl, bucketName, targetName), nil
}