package main

import (
	"log"
	"net/http"

	"github.com/sheershjindal/visitor-management-system/config"
	"github.com/sheershjindal/visitor-management-system/registry"
)

func main() {
	// Initialize app configurations and connections
	appConfig := config.Initialize()
	defer appConfig.Close()

	// Set up HTTP server
	mux := http.NewServeMux()

	// Initialize app registry with all modules and routes
	appRegistry := registry.NewAppRegistry(appConfig)
	appRegistry.RegisterAllRoutes(mux)

	// Start server
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
