package handlers

import (
	"encoding/json"
	"github.com/JackDaniells/pack-service/domain/contracts"
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
	packs := s.packService.Calculate(intItems)

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
