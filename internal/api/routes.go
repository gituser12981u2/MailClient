package api

import (
	"mailclient/internal/api/handlers"
	"mailclient/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authService *services.AuthService, emailService *services.EmailService) {
	api := router.Group("/api/v1")

	// Email routes
	emails := api.Group("/emails")
	{
		emails.GET("", handlers.GetEmails(emailService))
		emails.GET("/:id", handlers.GetEmail(emailService))
		emails.POST("", handlers.SendEmail(emailService))
	}
}
