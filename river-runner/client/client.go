package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/Condition17/fleet-services/river-runner/proto/river-runner"
	"google.golang.org/grpc"
)

func main() {
	//35.211.182.241:8080
	conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	fmt.Println("Connection:", conn)
	defer conn.Close()

	client := proto.NewRiverRunnerClient(conn)
	if resp, err := client.RunRiver(context.Background(), &proto.RunRequest{TestRunId: 279}); err != nil {
		fmt.Println("River run call error:", err)
		return
	} else {
		fmt.Println("River run call response:", resp)
		return
	}
}
