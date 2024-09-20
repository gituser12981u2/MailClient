// Package services provides business logic for the email client application.
package services

import (
	"mailclient/internal/models"
	"mailclient/internal/providers"
)

// EmailService handles email-related operations.
type EmailService struct {
	provider providers.EmailProvider
}

// NewEmailService creates a new EmailService instance.
func NewEmailService(provider providers.EmailProvider) *EmailService {
	return &EmailService{provider: provider}
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
	return s.provider.SendEmail(email)
}
