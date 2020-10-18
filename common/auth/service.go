package auth

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/client"
	microErrors "github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/server"

	"github.com/Condition17/fleet-services/common"
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
)

func ServiceAuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		userServiceClient := proto.NewUserService(common.GetFullExternalServiceName("user-service"), client.DefaultClient)
		res, err := userServiceClient.GetProfile(ctx, &proto.EmptyRequest{})
		if err != nil {
			return microErrors.Unauthorized(common.GetFullExternalServiceName("user-service"), fmt.Sprintf("%v", err))
		}

		return fn(context.WithValue(ctx, "User", res.User), req, resp)
	}
}
