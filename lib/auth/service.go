package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/micro/go-micro/v2/client"
	microErrors "github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	"github.com/Condition17/fleet-services/lib"
	userServiceAuth "github.com/Condition17/fleet-services/user-service/auth"
	userServiceProto "github.com/Condition17/fleet-services/user-service/proto/user-service"
)

func ServiceAuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		var userClaims *userServiceProto.User
		var err error

		userClaims, err = GetUserFromDecodedToken(ctx)

		if err != nil {
			return microErrors.Unauthorized(lib.GetFullExternalServiceName("user-service"), fmt.Sprintf("%v", err))
		}

		userServiceClient := userServiceProto.NewUserService(lib.GetFullExternalServiceName("user-service"), client.DefaultClient)
		_, err = userServiceClient.GetProfile(ctx, &userServiceProto.EmptyRequest{})
		if err != nil {
			// The user profile was not found in our database
			return microErrors.Unauthorized(lib.GetFullExternalServiceName("user-service"), fmt.Sprintf("%v", err))
		}

		// Add user data from token in request context
		user, _ := json.Marshal(userClaims)
		return fn(context.WithValue(ctx, "User", user), req, resp)
	}
}

func GetUserBytesFromContext(ctx context.Context) []byte {
	var usrBytes []byte
	usrBytes, _ = ctx.Value("User").([]byte)

	return usrBytes
}

func GetAuthorizationBytesFromContext(ctx context.Context) []byte {
	meta, _ := metadata.FromContext(ctx)

	return []byte(meta["Authorization"])
}

func GetTokenBytesFromContext(ctx context.Context) []byte {
	splitAuthToken := strings.Split(string(GetAuthorizationBytesFromContext(ctx)), "Bearer ")

	if len(splitAuthToken) < 2 {
		return nil
	}

	return []byte(splitAuthToken[1])
}

func GetUserFromDecodedToken(ctx context.Context) (*userServiceProto.User, error) {
	var tokenService *userServiceAuth.TokenService = &userServiceAuth.TokenService{}
	var tokenBytes []byte = GetTokenBytesFromContext(ctx)
	claims, err := tokenService.Decode(string(tokenBytes))

	if err != nil {
		return nil, err
	}

	return claims.User, nil
}
