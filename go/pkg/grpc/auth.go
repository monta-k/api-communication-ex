package grpc

import (
	"api-communication-ex/pkg/auth"
	"context"
	"errors"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
)

func Authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	if token != "token" {
		return nil, errors.New("unauthorized")
	}

	ctx = context.WithValue(ctx, auth.UserCtxKey, &auth.User{Name: "user"})

	return ctx, nil
}
