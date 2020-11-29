package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/Condition17/fleet-services/file-builder/proto/file-builder"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("35.224.241.90:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	fmt.Println("Connection:", conn)
	defer conn.Close()

	client := proto.NewFileBuilderClient(conn)
	if resp, err := client.TestCall(context.Background(), &proto.FileAssembleRequest{TestRunId: 1}); err != nil {
		fmt.Println("Assemble call error:", err)
		return
	} else {
		fmt.Println("Call response:", resp)
		return
	}
}
