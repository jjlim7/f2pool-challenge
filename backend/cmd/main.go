package main

import (
	"backend/internals/handlers"
	"backend/internals/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)

	// Define routes
	router.HandleFunc("/", handlers.RootHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	router.HandleFunc("/v1/tools/lookup", handlers.LookupHandler).Methods("GET")
	router.HandleFunc("/v1/tools/history", handlers.HistoryHandler).Methods("GET")
	router.HandleFunc("/v1/tools/validate", handlers.ValidateHandler).Methods("POST")

	// Prometheus metrics endpoint
	router.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server. Default to Port 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %v", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
