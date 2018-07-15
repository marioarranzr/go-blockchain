package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
	"time"
	"github.com/mario800ml/go-blockchain/internal/server"
	"github.com/mario800ml/go-blockchain/internal/handler"
)

// -----------------------------------------------------------------------------
// main
// -----------------------------------------------------------------------------
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := handler.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		handler.Blockchain = append(handler.Blockchain, genesisBlock)
	}()
	log.Fatal(server.Run())
}
