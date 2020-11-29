package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/Condition17/fleet-services/file-builder/proto/file-builder"
	"google.golang.org/grpc"
)

func main() {
	//35.211.182.241:8080
	conn, err := grpc.Dial("35.211.182.241:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	fmt.Println("Connection:", conn)
	defer conn.Close()

	client := proto.NewFileBuilderClient(conn)
	if resp, err := client.AssembleFile(context.Background(), &proto.FileAssembleRequest{TestRunId: 34}); err != nil {
		fmt.Println("Assemble call error:", err)
		return
	} else {
		fmt.Println("Call response:", resp)
		return
	}
}
