package main

import (
	"log"
	"os"

	"github.com/vaibhavhapani/Distributed-In-Memory-Key-Value-Store/internal/store"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatalf("Usage: store <port>")
	}

	port := args[0]
	capacity := store.DEFAULT_CAPACITY

	store.InitStoreServer(":"+port, uint32(capacity)) // blocking call
}
