package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	fileServicePb "github.com/Condition17/fleet-services/file-service/proto/file-service/grpc"
	resourceManagerPb "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service/grpc"
	"github.com/Condition17/fleet-services/river-runner/config"
	proto "github.com/Condition17/fleet-services/river-runner/proto/river-runner"
	runControllerPb "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

const riverImageName string = "cconache/river3:latest"
const runStateTopic = "test-run-state"

// GLOBALS -- todo: fix this
var (
	testRunStateTopic     *pubsub.Topic
	fileServiceClient     fileServicePb.FileServiceClient
	resourceManagerClient resourceManagerPb.ResourceManagerServiceClient
)

type riverRunnerServer struct {
	proto.UnimplementedRiverRunnerServer
}

func (s *riverRunnerServer) RunRiver(ctx context.Context, req *proto.RunRequest) (*proto.EmptyResponse, error) {
	testRunId := req.TestRunId
	// get file system details for test run
	response, err := resourceManagerClient.GetFileSystem(context.Background(), &resourceManagerPb.FileSystemSpec{TestRunId: testRunId})
	if err != nil {
		// TODO: return error response here to caller
		log.Printf("Error encountered while retrieving file system details for the provided test run (id: %v): %v\n", testRunId, err)
		return nil, errors.New(fmt.Sprintf("Error encountered while retrieving file system details for the provided test run (id: %v): %v\n", testRunId, err))
	}
	var fileSystemDetails *resourceManagerPb.FileSystem = response.FileSystem
	log.Println("Retrieved file system details:", fileSystemDetails)
	// get file details
	res, err := fileServiceClient.ReadFile(context.Background(), &fileServicePb.File{Id: fileSystemDetails.TestRun.FileId})
	if err != nil {
		// TODO: return data to the caller here
		log.Printf("Error encountered while retrieving file details for provided test run (id: %v): %v\n", testRunId, err)
		return nil, errors.New(fmt.Sprintf("Error encountered while retrieving file details for provided test run (id: %v): %v\n", testRunId, err))
	}
	var fileDetails *fileServicePb.File = res.File
	log.Println("Retrieved file details:", fileDetails)

	// create target directory to mount files to
	nfsIpAddr := fileSystemDetails.IP
	nfsFileSharePath := fmt.Sprintf(":/%s", fileSystemDetails.FileShareName)
	mountVolumePath := path.Join("/mnt/", fmt.Sprintf("testrun_%v", req.TestRunId))
	// Ensure mount directory is created and ignore any other issue
	_ = os.Mkdir(mountVolumePath, 0700)
	// --- mount volume
	fmt.Println("Mounting volume...")
	if err := syscall.Mount(nfsFileSharePath, mountVolumePath, "nfs", 0, fmt.Sprintf("nolock,addr=%s", nfsIpAddr)); err != nil {
		log.Printf("Syscall mount error: %v\n", err)
		return nil, errors.New(fmt.Sprintf("Syscall mount error: %v\n", err))
	}
	log.Println("Successfully mounted volume at ", mountVolumePath, "...")

	go func() {
		runFinishChan := make(chan bool)
		runErrorChan := make(chan error)

		go runRiverContainer(mountVolumePath,
			[]string{
				"-bp", fmt.Sprintf("%s/%s", mountVolumePath, fileDetails.Name),
				"-secondsBetweenStats", "2",
				"-arch", "x64",
				"-max", "1",
				"-outputType", "textual",
			}, runFinishChan, runErrorChan)

		select {
		case _ = <-runFinishChan:
			log.Println("Run successfully finished")
		case err := <-runErrorChan:
			// TODO: this case is not handle bellow
			log.Printf("River container run encountered error: %v", err)
		}

		// unmount volume
		fmt.Println("Trying unmount volume ", mountVolumePath)
		out, err := exec.Command("umount", "-l", mountVolumePath).Output()
		if err != nil {
			// TODO: handle this case
			log.Fatalf("Error unmounting fs: %s | Out: %s", err, out)
		}
		_ = os.Remove(mountVolumePath)
		fmt.Println("Successfully unmounted")

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
			// TODO: handle this case, too
			log.Println("Error encountered sending message to run controller service:", err)
		}
		log.Printf("Published message to '%v' topic. Message ID: %v\n", testRunStateTopic.String(), id)
	}()

	return &proto.EmptyResponse{}, nil
}

func runRiverContainer(volumePath string, cmdArgs []string, finishChan chan<- bool, errorChan chan<- error) {
	dockerCmd := fmt.Sprintf("docker run -v %s:/mount %s %s", volumePath, riverImageName, strings.Join(cmdArgs, " "))
	if err := exec.Command("/bin/sh", "-c", dockerCmd).Run(); err != nil {
		errorChan <- err
	}
	finishChan <- true
}

func main() {
	configs := config.GetConfig()

	// Server startup
	var serverAddr string = fmt.Sprintf("0.0.0.0:%v", configs.ServerPort)
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Server failed to listen on port ':%v'. Error encountered: %v\n", configs.ServerPort, err)
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	proto.RegisterRiverRunnerServer(grpcServer, &riverRunnerServer{})

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

	log.Printf("Starting GRPC server on localhost:%v\n", configs.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
