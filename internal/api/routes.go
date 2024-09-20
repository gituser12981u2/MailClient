// Package api provides the API setup and routing for the email client.
package api

import (
	"mailclient/internal/api/handlers"
	"mailclient/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes for the email client.
// It sets up the routes for email-related operations.
func SetupRoutes(router *gin.Engine, authService *services.AuthService, emailService *services.EmailService) {
	api := router.Group("/api/v1")

	// Email routes
	emails := api.Group("/emails")
	{
		emails.GET("", handlers.GetEmails(emailService))
		emails.GET("/:id", handlers.GetEmail(emailService))
		emails.POST("", handlers.SendEmail(emailService))
	}

	// TODO: Add routes for authentication, user management, and other features
}
