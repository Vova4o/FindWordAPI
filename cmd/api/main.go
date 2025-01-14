package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vova4o/findwordapi/config"
)

// @title Find a Word API
// @version 1.0
// @description This is a simple API for finding words in a list of words.
// @description It has a single endpoint that filters the list of words based on the letters and numbers provided in the input in cyrillic letters.

// @host http://159.89.17.9:8081
// @BasePath /api

type application struct {
	cfg config.Config
}

func main() {
	flag.Parse()

	app := application{
		cfg: config.NewConfig(),
	}

	// Initialize the router
	router := app.routes()

	// Combine host and port for the full address
	address := app.cfg.Host + app.cfg.Port

	// Create a new http.Server instance
	server := &http.Server{
		Addr:              address,
		Handler:           router,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second, // Set the ReadTimeout to 5 seconds
		ReadHeaderTimeout: 30 * time.Second, // Set the ReadHeaderTimeout to 2 seconds
		WriteTimeout:      30 * time.Second, // Set the WriteTimeout to 10 seconds
	}

	// Start the server
	fmt.Println("Starting server on", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
