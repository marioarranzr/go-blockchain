package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/marioarranzr/go-blockchain/internal/handler"
)

// -----------------------------------------------------------------------------
// Server
// -----------------------------------------------------------------------------
func Run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	if httpAddr == "" {
		httpAddr = "8080"
	}
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handler.HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handler.HandleWriteBlock).Methods("POST")
	return muxRouter
}
