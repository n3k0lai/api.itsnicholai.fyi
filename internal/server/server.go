package server

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"fyi/internal/handlers"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.HandleFunc("/about", handlers.About).Methods("GET")
	r.HandleFunc("/experience", handlers.Experience).Methods("GET")
	r.HandleFunc("/projects", handlers.Projects).Methods("GET")
	r.HandleFunc("/contact", handlers.Contact).Methods("GET")

	// Serve static files if the directory exists
	staticDir := "./static"
	if _, err := os.Stat(staticDir); !os.IsNotExist(err) {
		fileServer := http.FileServer(http.Dir(staticDir))
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	}

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	return c.Handler(r)
}
