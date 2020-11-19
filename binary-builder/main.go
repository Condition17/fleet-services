package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"path/filepath"
	"syscall"
	"os/exec"
)

func getMountCmd(os string) string {
	switch os {
	case "linux":
		return "mount"
	case "darwin":
		// mac os
		return "mount_nfs"
	default:
		return ""
	}
}

func getMountCmdArgs(os string) []string {
	switch os {
	case "darwin":
		return []string{"-o", "resvport"}
	default:
		return []string{}
	}
}

func main() {
	// config := config.GetConfig()

	// // Create google storage client
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

	wdPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(runtime.GOOS)
	addr := "10.24.195.50"
	source := ":/target"
	targetPath := fmt.Sprintf("%s%s", wdPath, "/mnt")

	if err := syscall.Mount(source, targetPath, "nfs", 0, fmt.Sprintf("nolock,addr=%s", addr)); err != nil {
		log.Fatalf("Syscall mount error: %v", err)
	}
	fmt.Println("NFS successfully mounted.")

	// try file creation in NFS
	f, err := os.OpenFile(filepath.Join(targetPath, "program_file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	fmt.Println("Trying to unmount")
	out, err := exec.Command("umount","-l", targetPath).Output()
	if err != nil {
		log.Fatalf("Error unmounting fs: %s | Out: %s", err, out)
	}
	fmt.Println("Successfully unmounted")
	fmt.Println(string(out[:]))
}
