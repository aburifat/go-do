package main

import (
	"log"
	"net/http"

	"github.com/aburifat/go-do/routes"
	"github.com/aburifat/go-do/storage"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()

	router := routes.SetupRouter(memoryStorage)

	log.Println("Server is running on http://localhost:8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
