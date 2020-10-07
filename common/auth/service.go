package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/micro/go-micro/v2/client"
	microErrors "github.com/micro/go-micro/v2/errors"
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
		// don't know why is uppercase
		token := meta["Token"]

		// Auth here
		userServiceClient := proto.NewUserService(common.GetFullExternalServiceName("user-service"), client.DefaultClient)
		if _, err := userServiceClient.ValidateToken(context.Background(), &proto.Token{Token: token}); err != nil {
			return microErrors.Unauthorized(common.GetFullExternalServiceName("user-service"), fmt.Sprintf("%v", err))
		}

		return fn(ctx, req, resp)
	}
}
