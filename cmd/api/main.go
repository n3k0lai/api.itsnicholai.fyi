package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"fyi/internal/server"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "6000" // Default port
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Invalid port %s: %v\n", portStr, err)
		os.Exit(1)
	}

	router := server.NewRouter()

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server starting on %s\n", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}
