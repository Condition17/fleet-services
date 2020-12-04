package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	resourceManagerPb "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	"github.com/Condition17/fleet-services/river-runner/config"
	runControllerPb "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"google.golang.org/grpc"
	"log"
	"os/exec"
	"strings"
)

const riverImageName string = "cconache/river3:latest"
const runStateTopic = "test-run-state"

// GLOBALS -- todo: fix this
var (
	testRunStateTopic     *pubsub.Topic
	fileServiceClient     fileServicePb.FileServiceClient
	resourceManagerClient resourceManagerPb.ResourceManagerServiceClient
)

func runRiverContainer(volumePath string, cmdArgs []string, finishChan chan<- bool, errorChan chan<- error) {
	dockerCmd := fmt.Sprintf("docker run -v %s:/mount %s %s", volumePath, riverImageName, strings.Join(cmdArgs, " "))
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

	// Create pubsub client
	pubSubClient, err := pubsub.NewClient(context.Background(), configs.GoogleProjectID)
	if err != nil {
		log.Fatalf("Error on pubsub.NewClient: %v", err)
	}

	// Get test run state topic
	testRunStateTopic = pubSubClient.Topic(runStateTopic)

	// Initialize clients of external services
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

	// should mount volume at a certain path
	// TODO: replace this
	mountVolumePath := "/Users/cristian_conache/go/src/fleet/river-runner/mount"


	// Start the river container here
	runFinishChan := make(chan bool)
	runErrorChan := make(chan error)

	go runRiverContainer(mountVolumePath,
		[]string{
			"-bp", fmt.Sprintf("/mount/%s", fileDetails.Name),
			"-secondsBetweenStats", "2",
			"-arch", "x64",
			"-max", "1",
			"-outputType", "textual",
		}, runFinishChan, runErrorChan)

	select {
	case _ = <-runFinishChan:
		log.Println("Run successfully finished")
		// put data on run state queue
	case err := <-runErrorChan:
		log.Printf("River container run encountered error: %v", err)
		// put data on run state queue
	}

	// construct the notification message
	eventData, _ := json.Marshal(&runControllerPb.RiverRunFinishedEventData{TestRunId: testRunId})
	// send notification
	msg, _ := json.Marshal(&runControllerPb.Event{
			Type: "test-run.finished",
			Meta: &runControllerPb.EventMetadata{Authorization: []byte("")},
			Data: eventData,
	})
	result := testRunStateTopic.Publish(context.Background(), &pubsub.Message{Data: msg})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Println("Error encountered sending message to run controller service:", err)
		return
	}
	log.Printf("Published message to '%v' topic. Message ID: %v\n", testRunStateTopic.String(), id)
}
