package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/Condition17/fleet-services/river-runner/config"
	"github.com/Condition17/fleet-services/river-runner/handler"
	proto "github.com/Condition17/fleet-services/river-runner/proto/river-runner"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var lis net.Listener
	var conn *grpc.ClientConn
	var pubSubClient *pubsub.Client
	var err error
	configs := config.GetConfig()

	// Server startup
	if lis, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", configs.ServerPort)); err != nil {
		log.Fatalf("Server failed to initiate listener on port ':%v': %v\n", configs.ServerPort, err)
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	// Create connection to Fleet micro services proxy
	if conn, err = grpc.Dial(configs.FleetServicesGrpcProxyUrl, grpc.WithInsecure()); err != nil {
		log.Fatalf("Error encountered while creating connection to Fleet micro services GRPC proxy: %v", err)
	}
	log.Println("Connection to fleet services GRPC proxy successfully initiated")
	defer conn.Close()

	// Create Google Pub Sub client
	if 	pubSubClient, err = pubsub.NewClient(context.Background(), configs.GoogleProjectID); err != nil {
		log.Fatalf("Error creating Google Pub Sub client: %v", err)
	}

	proto.RegisterRiverRunnerServer(grpcServer, handler.NewHandler(conn, pubSubClient))

	log.Printf("Starting GRPC server on localhost:%v\n", configs.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
