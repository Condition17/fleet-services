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
		meta, _ := metadata.FromContext(ctx)
		var isServiceCaller bool = meta["Host"] == "" || meta["Method"] == ""
		fmt.Printf("Meta: %v - %v\n", meta, isServiceCaller)
		var reqCtx = context.WithValue(ctx, "serviceCaller", isServiceCaller)

		if isServiceCaller {
			return fn(reqCtx, req, resp)
		}

		userServiceClient := userServiceProto.NewUserService(lib.GetFullExternalServiceName("user-service"), client.DefaultClient)

		if _, err := userServiceClient.GetProfile(ctx, &userServiceProto.EmptyRequest{}); err != nil {
			// The user profile was not found in our database
			return microErrors.Unauthorized(lib.GetFullExternalServiceName("user-service"), fmt.Sprintf("%v", err))
		}

		return fn(reqCtx, req, resp)
	}
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

func GetUserBytesFromDecodedToken(ctx context.Context) []byte {
	var userClaims *userServiceProto.User
	userClaims, err := GetUserFromDecodedToken(ctx)

	if err != nil {
		fmt.Printf("Error encountered while extracting user from decoded token: %v", err)
		return nil
	}

	usrBytes, _ := json.Marshal(userClaims)
	return usrBytes
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
