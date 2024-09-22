// Package handlers provides HTTP request handlers for the email client API.
package handlers

import (
	"net/http"
	"strconv"

	"mailclient/internal/models"
	"mailclient/internal/services"

	"github.com/gin-gonic/gin"
)

// GetEmails returns a handler function for retrieving emails.
// It accepts query parameters for folder, limit, and offset.
func GetEmails(emailService *services.EmailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		folder := c.DefaultQuery("folder", "inbox")
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		emails, err := emailService.GetEmails(folder, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, emails)
	}
}

// GetEmail returns a handler function for retrieving a single email by its ID.
func GetEmail(emailService *services.EmailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		email, err := emailService.GetEmail(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, email)
	}
}

// SendEmail returns a handler function for sending a new email.
func SendEmail(emailService *services.EmailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var email models.Email
		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := emailService.SendEmail(&email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	}
}
