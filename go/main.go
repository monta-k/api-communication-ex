package main

import (
	oapicodegen "api-communication-ex/oapi-codegen"
	"api-communication-ex/oapi-codegen/adapters"
	"log"
	"net/http"
)

func main() {
	oapiCodegenServer := oapicodegen.NewOAPICodeGenServer()
	r := http.NewServeMux()
	h := adapters.HandlerFromMux(oapiCodegenServer, r)
	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
