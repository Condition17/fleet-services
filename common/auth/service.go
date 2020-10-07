package auth

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	"github.com/Condition17/fleet-services/common"
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
)

func ServiceAuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		log.Println()
		// don't know why is uppercase
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		userServiceClient := proto.NewUserService(common.GetFullExternalServiceName("user"), client.DefaultClient)
		_, err := userServiceClient.ValidateToken(context.Background(), &proto.Token{Token: token})

		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)

		return err
	}
}
