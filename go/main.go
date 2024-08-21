package main

import (
	"api-communication-ex/gqlgen/generated"
	"api-communication-ex/gqlgen/graph"
	oapicodegen "api-communication-ex/oapi-codegen"
	"api-communication-ex/oapi-codegen/adapters"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
