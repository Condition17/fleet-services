package main

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os/exec"
	"strings"
	"context"
	"github.com/Condition17/fleet-services/river-runner/config"
	resourceManagerPb "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"

)

const riverImageName string = "cconache/river3:latest"

// GLOBALS -- todo: fix this
var (
	testRunStateTopic     *pubsub.Topic
	fileServiceClient     fileServicePb.FileServiceClient
	resourceManagerClient resourceManagerPb.ResourceManagerServiceClient
)

func runRiverContainer(volumePath string, cmdArgs []string, finishChan chan<- bool, errorChan chan<- error) {
	dockerCmd := fmt.Sprintf("sudo docker run -v %s:/mount %s %s", volumePath, riverImageName, strings.Join(cmdArgs, " "))
	if err := exec.Command("/bin/sh", "-c", dockerCmd).Run(); err != nil {
		errorChan <- err
	}
	finishChan <- true
}

func main() {
	configs := config.GetConfig()

	testRunId := uint32(279)
	// Server startup
	conn, err := grpc.Dial(configs.FleetServicesGrpcProxyUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error encountered while creating connection to fleet services grpc proxy: %v", err)
	}
	log.Println("Connection to fleet services GRPC proxy initiated")
	defer conn.Close()

	resourceManagerClient = resourceManagerPb.NewResourceManagerServiceClient(conn)
	fileServiceClient = fileServicePb.NewFileServiceClient(conn)

	// get file system details for test run
	response, err := resourceManagerClient.GetFileSystem(context.Background(), &resourceManagerPb.FileSystemSpec{TestRunId: testRunId})
	if err != nil {
		// TODO: return error response here to caller
		log.Printf("Error encountered while retrieving file system details for the provided test run (id: %v): %v\n", testRunId, err)
		return
	}
	var fileSystemDetails *resourceManagerPb.FileSystem = response.FileSystem

	log.Println("Retrieved file system details:", fileSystemDetails)
	// get file details
	res, err := fileServiceClient.ReadFile(context.Background(), &fileServicePb.File{Id: fileSystemDetails.TestRun.FileId})
	if err != nil {
		// TODO: return data to the caller here
		log.Printf("Error encountered while retrieving file details for provided test run (id: %v): %v\n", testRunId, err)
		return
	}
	var fileDetails *fileServicePb.File = res.File
	log.Println("Retrieved file details:", fileDetails)

	//volumePath := "/Users/cristian_conache/Workspace/learning/river/mount"
	//runFinishChan := make(chan bool)
	//runErrorChan := make(chan error)
	//
	//go runRiverContainer(volumePath,
	//	[]string{
	//		"-bp", fmt.Sprintf("/mount/%s", fileDetails.Name),
	//		"-secondsBetweenStats", "2",
	//		"-arch", "x64",
	//		"-max", "1",
	//		"-outputType", "textual",
	//	}, runFinishChan, runErrorChan)
	//
	//select {
	//case _ = <-runFinishChan:
	//	log.Println("Run successfully finished")
	//case err := <-runErrorChan:
	//	log.Printf("River container run encountered error: %v", err)
	//}
}
