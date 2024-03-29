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
	r.Use(publicEndpointMiddleware)
	r.HandleFunc("/calculate", packHandler.Calculate).Methods(http.MethodGet)
	r.HandleFunc("/packs", packHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/packs", packHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/packs/{pack}", packHandler.Remove).Methods(http.MethodDelete)
	r.HandleFunc("/packs/list", packHandler.AddList).Methods(http.MethodPost)
	r.HandleFunc("/packs/list", packHandler.UpdateList).Methods(http.MethodPut)
	r.HandleFunc("/packs/list/remove", packHandler.RemoveList).Methods(http.MethodPost)

	r.HandleFunc("/packs", defaultOptionsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/packs/{pack}", defaultOptionsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/packs/list", defaultOptionsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/packs/list/remove", defaultOptionsHandler).Methods(http.MethodOptions)
	return r
}

func publicEndpointMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}

func defaultOptionsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
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
