package main

import (
	"context"
	"fmt"
	"log"

	binaryBuilderProto "github.com/Condition17/fleet-services/binary-builder/proto/binary-builder"
	"github.com/Condition17/fleet-services/run-controller-service/config"
	handler "github.com/Condition17/fleet-services/run-controller-service/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/broker/googlepubsub/v2"
	"google.golang.org/grpc"
)

const topic string = "test-run-state"

func main() {
	configs := config.GetConfig()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.run-controller-service"),
		micro.Broker(googlepubsub.NewBroker(googlepubsub.ProjectID(configs.GoogleProjectID))),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Get the message broker instance
	msgBroker := service.Server().Options().Broker
	if err := msgBroker.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe run state topic
	_, err := msgBroker.Subscribe(topic, handler.NewHandler(service))

	if err != nil {
		log.Fatal(err)
	}

	// -- remove this block
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	fmt.Println("Connection:", conn)
	defer conn.Close()

	client := binaryBuilderProto.NewBinaryBuilderClient(conn)
	if resp, err := client.Hello(context.Background(), &binaryBuilderProto.EmptyMessage{}); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Call response:", resp)
		return
	}
	// ---

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
