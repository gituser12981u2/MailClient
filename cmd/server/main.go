// Package main is the entry point for the email client application.
// It sets up the configuration, initializes services, and starts the HTTP server.
package main

import (
	"log"

	"mailclient/internal/api"
	"mailclient/internal/config"
	"mailclient/internal/providers"
	"mailclient/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to create email server: %v", err)
	}

	// Initialize email provider
	provider, err := providers.New(cfg.ProviderConfig.Type, cfg.ProviderConfig.Config)
	if err != nil {
		log.Fatalf("Failed to initialize email provider: %v", err)
	}

	// Initialize services
	authService := services.NewAuthService(cfg.JWTSecret)
	emailService := services.NewEmailService(provider, cfg.ProviderConfig.Config["username"])

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	api.SetupRoutes(router, authService, emailService)

	// Start server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
