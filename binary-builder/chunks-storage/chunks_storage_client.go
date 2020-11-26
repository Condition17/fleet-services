package chunksStorage

import (
	"cloud.google.com/go/storage"
	"context"
)
type Client struct {
	bucket *storage.BucketHandle
}

func NewChunksStorageClient(googleProjectID string, bucketName string) (*Client, error){
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &Client{bucket: client.Bucket(bucketName).UserProject(googleProjectID)}, nil
}

func (c *Client) GetObjectReader(objectName string) (*storage.Reader, error){
	return c.bucket.Object(objectName).NewReader(context.Background())
}