package main

import (
	"github.com/JackDaniells/pack-service/api"
	"github.com/JackDaniells/pack-service/domain/handlers"
	"github.com/JackDaniells/pack-service/domain/repository"
	"github.com/JackDaniells/pack-service/domain/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("starting pack service application")

	packRepository := repository.NewPackRepository()
	packService := service.NewPackService(packRepository)
	packHandler := handlers.NewPackHandler(packService)

	apiPort := "8080"
	server := api.NewServer(apiPort, packHandler)

	// Serve api routes
	server.Serve()
	defer server.GracefulShutdown()
	log.Println("API server started, listening on port: ", apiPort)

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("shutting application down")
}
