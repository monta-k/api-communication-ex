package connect

import (
	"api-communication-ex/pkg/auth"
	"context"
	"errors"

	"connectrpc.com/connect"
)

const authorizationHeader = "authorization"

func NewAuthClientInterceptor(token string) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				req.Header().Set(authorizationHeader, token)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func NewAuthServerInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if !req.Spec().IsClient {
				if token := req.Header().Get(authorizationHeader); token != "token" {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthorized"))
				}
				ctx = context.WithValue(ctx, auth.UserCtxKey, &auth.User{Name: "user"})
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
