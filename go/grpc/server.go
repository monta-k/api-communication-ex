package grpc

import (
	"context"
	"fmt"

	hellopb "api-communication-ex/grpc/pkg/grpc"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", req.GetName()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}
