package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"your-module-path/internal/handlers"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.HandleFunc("/about", handlers.About).Methods("GET")
	r.HandleFunc("/experience", handlers.Experience).Methods("GET")
	r.HandleFunc("/projects", handlers.Projects).Methods("GET")
	r.HandleFunc("/contact", handlers.Contact).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	return c.Handler(r)
}
