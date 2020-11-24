package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	"google.golang.org/grpc"
)

func main() {
	// -- server startup

	// -- handling request
	// get file details
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	fmt.Println("Connection:", conn)
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	if fileDetails, err := client.ReadFile(context.Background(), &pb.File{Id: "ea223793-5b74-49df-8806-b1e5c3d4e064"}); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("File details:", fileDetails)
		return
	}

	// if fileDetails, err := client.ReadFile(context.Background(), &fileServiceProto.File{Id: "file:ea223793-5b74-49df-8806-b1e5c3d4e064:uploadedChunksCount"}); err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// } else {
	// 	fmt.Println("File details:", fileDetails)
	// 	return
	// }
	// var fileID string = "ea223793-5b74-49df-8806-b1e5c3d4e064"
	// if fileDetails == nil {
	// 	fmt.Println("File not found")
	// 	return
	// }

	// if err != nil {
	// 	fmt.Printf("Error encountered while retrieving file details: %v", fileID)
	// 	return
	// }

	// -- spawn goroutine
	// -- request handling end
	return

	// --------------------------------------
	// -- EXPERIMENTS --

	// Create google storage client
	// client, err := storage.NewClient(context.Background())
	// if err != nil {
	// 	log.Fatalf("Failed to create Google Storage Client: %v", err)
	// }
	// defer client.Close()

	// bucketName := "fleet-files-chunks"

	// // Create bucket instance
	// bucket := client.Bucket(bucketName).UserProject(config.GoogleProjectID)
	// attrs, errs := bucket.Attrs(context.Background())
	// fmt.Printf("bucket attrs: %v - error: %v", attrs, errs)
	// objectName := "1a510b80dacb2e5251df15becafd41619ebeb7e9eb1c97ce0de9cfa1832cf5d4"
	// reader, err := bucket.Object(objectName).NewReader(context.Background())
	// if err != nil {
	// 	log.Fatalf("Object(%q).NewReader: %v", objectName, err)
	// }
	// defer reader.Close()

	// data, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	fmt.Errorf("ioutil.ReadAll: %v", err)
	// }
	// fmt.Printf("Blob %v downloaded: %v\n", objectName, data)

	// syscall.Mount("10.252.184.154:/target", "./mnt", "nfs")

	// wdPath, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(runtime.GOOS)
	// addr := "10.24.195.50"
	// source := ":/target"
	// targetPath := fmt.Sprintf("%s%s", wdPath, "/mnt")

	// if err := syscall.Mount(source, targetPath, "nfs", 0, fmt.Sprintf("nolock,addr=%s", addr)); err != nil {
	// 	log.Fatalf("Syscall mount error: %v", err)
	// }
	// fmt.Println("NFS successfully mounted.")

	// // try file creation in NFS
	// f, err := os.OpenFile(filepath.Join(targetPath, "program_file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f.Close()

	// fmt.Println("Trying to unmount")
	// out, err := exec.Command("umount", "-l", targetPath).Output()
	// if err != nil {
	// 	log.Fatalf("Error unmounting fs: %s | Out: %s", err, out)
	// }
	// fmt.Println("Successfully unmounted")
	// fmt.Println(string(out[:]))

	// file, _ := os.OpenFile("test.zip", os.O_CREATE|os.O_RDWR, 0666)
	// _, err := file.Seek(100, 0)
	// if err != nil {
	// 	fmt.Println("Error while seeking:", err)
	// 	return
	// }

	// if _, err := file.Write([]byte("this is a test")); err != nil {
	// 	fmt.Println("Error writing file:", err)
	// }

	// if bytes, err := ioutil.ReadFile("test.zip"); err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// } else {
	// 	fmt.Println(bytes)
	// 	fmt.Printf("File content: '%s'\n", bytes)
	// }
	// defer file.Close()
}
