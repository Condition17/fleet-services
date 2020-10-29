package main

import (
	"context"
	"fmt"

	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/micro/go-micro/v2"
)

func main() {
	var service micro.Service
	var baseHandler baseservice.BaseHandler

	service = micro.NewService(
		micro.Name("trial-service"),
		micro.Version("latest"),
	)

	// initialise test handler
	service.Init()

	baseHandler = baseservice.NewBaseHandler(service)

	// Test getTokenBytesFromContext
	// var ctx context.Context = metadata.Set(context.Background(), "Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjo3LCJuYW1lIjoiR2lnaWVsIFRlc3QiLCJjb21wYW55IjoiQWxwaGFiZXQgSW5jLiIsImVtYWlsIjoiZ2lnZWwuZW1haWxAZ21haWwuY29tIn0sImV4cCI6MTYzNDE1NDE5MiwiaXNzIjoiZ28ubWljcm8uYXBpLnVzZXItc2VydmljZSJ9.93BX3wcbjeTYszsR8UvpFLG1v1XS8GzQ7n7RzQcF3rU")
	// var vals []byte
	// vals = ctx.Value("Token").([]byte)
	// fmt.Print(vals)
	// fmt.Print(auth.GetTokenBytesFromContext(ctx))
	baseHandler.SendDataToWssQueue(context.Background(), []byte("this is from lib"))
	fmt.Println("Message sent, theoretically")
}
