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
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to create email server: %v", err)
	}

	provider, err := providers.New(cfg.Provider, cfg.ProviderConfig)
	if err != nil {
		log.Fatalf("Failed to initialize email provider: %v", err)
	}

	// Initialize services
	authService := services.NewAuthService(cfg.JWTSecret)
	emailService := services.NewEmailService(provider)

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
