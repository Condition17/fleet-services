package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/Condition17/fleet-services/binary-builder/proto/file-builder"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	fmt.Println("Connection:", conn)
	defer conn.Close()

	client := proto.NewBinaryBuilderClient(conn)
	if resp, err := client.Hello(context.Background(), &proto.EmptyMessage{}); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Call response:", resp)
		return
	}
}
