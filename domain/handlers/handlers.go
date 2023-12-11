package handlers

import (
	"encoding/json"
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/JackDaniells/pack-service/domain/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type packHandler struct {
	packService contracts.PackService
}

func NewPackHandler(packService contracts.PackService) *packHandler {
	return &packHandler{
		packService: packService,
	}
}

func (s *packHandler) Calculate(response http.ResponseWriter, request *http.Request) {
	items := request.URL.Query().Get("items")

	intItems, err := strconv.Atoi(items)
	if err != nil {
		log.Println("error when parse items: ", err)
		http.Error(response, "error when parse items", http.StatusBadRequest)
		return
	}
	packs, err := s.packService.Calculate(intItems)
	if err != nil {
		log.Println("error when calculate pack items: ", err)
		http.Error(response, "error when calculate pack items: "+err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(packs)
	if err != nil {
		log.Println("error when marshal response: ", err)
		http.Error(response, "error when marshal response", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(bytes)
	if err != nil {
		log.Println("error when write response: ", err)
		http.Error(response, "error when write response", http.StatusInternalServerError)
		return
	}
}

func (s *packHandler) GetAll(response http.ResponseWriter, request *http.Request) {

	packs, err := s.packService.GetAll()
	if err != nil {
		log.Println("error when create pack: ", err)
		http.Error(response, "error when create pack: "+err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(packs)
	if err != nil {
		log.Println("error when marshal response: ", err)
		http.Error(response, "error when marshal response", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(bytes)
	if err != nil {
		log.Println("error when write response: ", err)
		http.Error(response, "error when write response", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (s *packHandler) Create(response http.ResponseWriter, request *http.Request) {

	var packRequest entity.Pack

	err := json.NewDecoder(request.Body).Decode(&packRequest)
	if err != nil {
		log.Println("error when decode request body: ", err)
		http.Error(response, "error when decode request body", http.StatusBadRequest)
		return
	}

	err = s.packService.Create(packRequest.Size)
	if err != nil {
		log.Println("error when create pack: ", err)
		http.Error(response, "error when create pack: "+err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (s *packHandler) Remove(response http.ResponseWriter, request *http.Request) {
	pack := mux.Vars(request)["pack"]

	IntPack, err := strconv.Atoi(pack)
	if err != nil {
		log.Println("error when parse pack: ", err)
		http.Error(response, "error when parse pack", http.StatusBadRequest)
		return
	}

	err = s.packService.Remove(IntPack)
	if err != nil {
		log.Println("error when remove pack: ", err)
		http.Error(response, "error when remove pack: "+err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusOK)
}
