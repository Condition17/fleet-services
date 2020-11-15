package main

import (
	"context"
	"fmt"

	"github.com/Condition17/fleet-services/lib/auth"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
)

func main() {
	var service micro.Service
	// var baseHandler baseservice.BaseHandler

	service = micro.NewService(
		micro.Name("trial-service"),
		micro.Version("latest"),
	)

	// initialise test handler
	service.Init()

	// baseHandler = baseservice.NewBaseHandler(service)

	// Test getTokenBytesFromContext
	var ctx context.Context = metadata.Set(context.Background(), "Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoxMiwibmFtZSI6InRzdDEiLCJjb21wYW55IjoidHN0MS5zcmwiLCJlbWFpbCI6InRzdDFAZGV2LmZsZWV0In0sImV4cCI6MTYzNTc5NzA4OCwiaXNzIjoiZ28ubWljcm8uYXBpLnVzZXItc2VydmljZSJ9.L9QKsNQzIeN3v-ov3jzYQjj3F5nmXIDrFKSYNArlKZw")
	var ctx1 context.Context = context.WithValue(context.Background(), "User", []byte("{test}"))
	// var vals []byte
	// vals = ctx.Value("Token").([]byte)
	// fmt.Print(vals)
	fmt.Println(string(auth.GetUserBytesFromDecodedToken(ctx)))
	fmt.Println(ctx1.Value("User"))
	fmt.Println("User bytes:", string(auth.GetUserBytesFromContext(ctx1)), "length:", len(auth.GetUserBytesFromContext(ctx1)))
	// baseHandler.SendDataToWssQueue(context.Background(), []byte("this is from lib"))
	// fmt.Println("Message sent, theoretically")
}
