package handlers

import (
	"net/http"
	"strconv"

	"mailclient/internal/models"
	"mailclient/internal/services"

	"github.com/gin-gonic/gin"
)

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

func SendEmail(emailService *services.EmailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var email models.Email
		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := emailService.SendEmail(&email); err != nil {
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
