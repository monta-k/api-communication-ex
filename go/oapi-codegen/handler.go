package oapicodegen

import (
	"api-communication-ex/oapi-codegen/adapters"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type OAPICodeGenServer struct{}

func NewOAPICodeGenServer() *OAPICodeGenServer {
	return &OAPICodeGenServer{}
}

func (OAPICodeGenServer) ListPets(w http.ResponseWriter, r *http.Request, params adapters.ListPetsParams) {
	fmt.Println("params", params)
	resp := adapters.Pets{
		adapters.Pet{
			Id:   1,
			Name: "Dog",
			Tag:  nil,
		},
		adapters.Pet{
			Id:   2,
			Name: "Cat",
			Tag:  nil,
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (OAPICodeGenServer) CreatePets(w http.ResponseWriter, r *http.Request) {
	var requestBody adapters.CreatePetsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	resp := adapters.Pet{
		Id:   requestBody.Id,
		Name: requestBody.Name,
		Tag:  requestBody.Tag,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (OAPICodeGenServer) ShowPetById(w http.ResponseWriter, r *http.Request, petId string) {
	id, err := strconv.ParseInt(petId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid petId", http.StatusBadRequest)
		return
	}
	fmt.Println("petId", id)
	resp := adapters.Pet{
		Id:   id,
		Name: "Dog",
		Tag:  nil,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
