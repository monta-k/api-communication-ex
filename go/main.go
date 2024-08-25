package main

import (
	"api-communication-ex/gqlgen/generated"
	"api-communication-ex/gqlgen/graph"
	oapicodegen "api-communication-ex/oapi-codegen"
	"api-communication-ex/oapi-codegen/adapters"
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

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	r := http.NewServeMux()

	// GraphQL server setup
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	r.Handle("/graphql", srv)
	r.Handle("/graphql-playground", playground.Handler("GraphQL playground", "/graphql"))

	// REST server setup
	oapiCodegenServer := oapicodegen.NewOAPICodeGenServer()
	h := adapters.HandlerFromMux(oapiCodegenServer, r)

	httpServer := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	grpcServer := grpc.NewServer()
	hellopb.RegisterGreetingServiceServer(grpcServer, mygrpcserver.NewMyServer())
	reflection.Register(grpcServer)

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
