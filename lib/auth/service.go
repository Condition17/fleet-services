package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/micro/go-micro/v2/client"
	microErrors "github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	"github.com/Condition17/fleet-services/lib"
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
)

func ServiceAuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		// Auth here
		userServiceClient := proto.NewUserService(lib.GetFullExternalServiceName("user-service"), client.DefaultClient)
		res, err := userServiceClient.GetProfile(ctx, &proto.EmptyRequest{})
		if err != nil {
			return microErrors.Unauthorized(lib.GetFullExternalServiceName("user-service"), fmt.Sprintf("%v", err))
		}

		user, _ := json.Marshal(res.User)
		return fn(context.WithValue(ctx, "User", user), req, resp)
	}
}

func GetUserBytesFromContext(ctx context.Context) []byte {
	var usrBytes []byte
	usrBytes, _ = ctx.Value("User").([]byte)

	return usrBytes
}

func GetTokenBytesFromContext(ctx context.Context) []byte {
	var tokenBytes []byte
	meta, _ := metadata.FromContext(ctx)
	tokenBytes = []byte(meta["Token"])

	return tokenBytes
}

func GetUserFromContext(ctx context.Context) (*proto.User, error) {
	var user *proto.User
	if err := json.Unmarshal(GetUserBytesFromContext(ctx), &user); err != nil {
		return nil, err
	}

	return user, nil
}
