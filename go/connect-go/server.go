package connecgo

import (
	greetv1 "api-communication-ex/connect-go/gen/greet/v1"
	"api-communication-ex/connect-go/gen/greet/v1/greetv1connect"
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
)

type GreetServer struct {
	greetv1connect.UnimplementedGreetServiceHandler
}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func NewGreetServer() *GreetServer {
	return &GreetServer{}
}
