package api

import (
	"context"
	"fmt"
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	wait       time.Duration
}

func NewMuxRouter(packHandler contracts.PackHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pack/calculate", packHandler.Calculate).Methods(http.MethodGet)
	return r
}

func NewServer(apiPort string, portHandler contracts.PackHandler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%s", apiPort),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			Handler:      NewMuxRouter(portHandler),
		},
		wait: 10 * time.Second,
	}
}

func (s *Server) Serve() {
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
	}()
}

func (s *Server) GracefulShutdown() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), s.wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("error to shutdown server: %v\n", err)
	}
}
