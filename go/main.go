package main

import (
	connecgo "api-communication-ex/connect-go"
	"api-communication-ex/connect-go/gen/greet/v1/greetv1connect"
	"api-communication-ex/gqlgen/generated"
	"api-communication-ex/gqlgen/graph"
	oapicodegen "api-communication-ex/oapi-codegen"
	"api-communication-ex/oapi-codegen/adapters"
	"api-communication-ex/pkg/auth"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mygrpcserver "api-communication-ex/grpc"
	hellopb "api-communication-ex/grpc/pkg/grpc"

	mygrpc "api-communication-ex/pkg/grpc"

	myconnect "api-communication-ex/pkg/connect"

	"connectrpc.com/connect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {
	r := http.NewServeMux()

	// GraphQL server setup
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	r.Handle("/graphql", auth.AuthMiddleware(srv))
	r.Handle("/graphql-playground", playground.Handler("GraphQL playground", "/graphql"))

	// connect-go server setup
	interceptors := connect.WithInterceptors(myconnect.NewAuthServerInterceptor())
	greetServer := connecgo.NewGreetServer()
	path, handler := greetv1connect.NewGreetServiceHandler(greetServer, interceptors)
	r.Handle(path, handler)

	// REST server setup
	oapiCodegenServer := oapicodegen.NewOAPICodeGenServer()
	h := adapters.HandlerWithOptions(oapiCodegenServer, adapters.StdHTTPServerOptions{
		BaseRouter: r,
		Middlewares: []adapters.MiddlewareFunc{
			auth.AuthMiddleware,
		},
	})

	h2cHandler := h2c.NewHandler(h, &http2.Server{})
	httpServer := &http.Server{
		Handler: h2cHandler,
		Addr:    "0.0.0.0:8080",
	}

	// gRPC server setup
	authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(mygrpc.Authenticate))
	grpcServer := grpc.NewServer(authInterceptor)
	hellopb.RegisterGreetingServiceServer(grpcServer, mygrpcserver.NewMyServer())

	// Start gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Starting gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Channel to listen for interrupt or terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start HTTP server in a goroutine
	go func() {
		log.Fatal(httpServer.ListenAndServe())
	}()
	log.Println("Starting HTTP server on :8080")

	// Wait for interrupt signal
	<-stop

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}

	// Stop gRPC server
	grpcServer.GracefulStop()

	log.Println("Servers gracefully stopped")
}
