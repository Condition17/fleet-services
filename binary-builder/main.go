package main

import (
	"github.com/Condition17/fleet-services/binary-builder/config"
	"context"
	"log"
	"io/ioutil"
	"fmt"

	"cloud.google.com/go/storage"
)

func main() {
	config := config.GetConfig()

	// Create google storage client
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Google Storage Client: %v", err)
	}
	defer client.Close()

	bucketName := "fleet-files-chunks"

	// Create bucket instance
	bucket := client.Bucket(bucketName).UserProject(config.GoogleProjectID)
	attrs, errs := bucket.Attrs(context.Background())
	fmt.Printf("bucket attrs: %v - error: %v", attrs, errs)
	objectName := "1a510b80dacb2e5251df15becafd41619ebeb7e9eb1c97ce0de9cfa1832cf5d4"
	reader, err := bucket.Object(objectName).NewReader(context.Background())
	if err != nil {
		log.Fatalf("Object(%q).NewReader: %v", objectName, err)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	fmt.Printf("Blob %v downloaded: %v\n", objectName, data)
}
