// Package services provides business logic for the email client application.
package services

import (
	"fmt"
	"time"

	"mailclient/internal/models"
	"mailclient/internal/providers"

	"github.com/google/uuid"
)

// EmailService handles email-related operations.
type EmailService struct {
	provider    providers.EmailProvider
	senderEmail string
}

// NewEmailService creates a new EmailService instance.
func NewEmailService(provider providers.EmailProvider, senderEmail string) *EmailService {
	return &EmailService{
		provider:    provider,
		senderEmail: senderEmail,
	}
}

// GetEmails retrieves emails from the specified folder with pagination.
func (s *EmailService) GetEmails(folder string, limit, offset int) ([]models.Email, error) {
	return s.provider.GetEmails(folder, limit, offset)
}

// GetEmail retrieves a single email by its ID.
func (s *EmailService) GetEmail(id string) (*models.Email, error) {
	return s.provider.GetEmail(id)
}

// SendEmail sends a new email.
func (s *EmailService) SendEmail(email *models.Email) error {
	// Set the From field if it's not already set
	if email.From == "" {
		email.From = s.senderEmail
	}

	// Set other fields if they're not set
	if email.Date.IsZero() {
		email.Date = time.Now()
	}
	if email.MessageID == "" {
		email.MessageID = fmt.Sprintf("<%s@mailclient>", uuid.New().String())
	}

	return s.provider.SendEmail(email)
}
