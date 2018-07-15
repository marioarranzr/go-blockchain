package server

import (
	"os"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/mario800ml/go-blockchain/internal/handler"
)

// -----------------------------------------------------------------------------
// Server
// -----------------------------------------------------------------------------
func Run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", os.Getenv("ADDR"))
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