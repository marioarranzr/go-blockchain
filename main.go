package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/marioarranzr/go-blockchain/internal/handler"
	"github.com/marioarranzr/go-blockchain/internal/server"
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
		genesisBlock := handler.Block{Timestamp: t.String()}
		spew.Dump(genesisBlock)
		handler.Blockchain = append(handler.Blockchain, genesisBlock)
	}()
	log.Fatal(server.Run())
}
