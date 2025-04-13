package main

import (
	"log"
	"net/http"
	"os"

	"fyi/internal/server"
)

func main() {
	router := server.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
