package services

import (
	"mailclient/internal/models"
	"mailclient/internal/providers"
)

type EmailService struct {
	provider providers.EmailProvider
}

func NewEmailService(provider providers.EmailProvider) *EmailService {
	return &EmailService{provider: provider}
}

func (s *EmailService) GetEmails(folder string, limit, offset int) ([]models.Email, error) {
	return s.provider.GetEmails(folder, limit, offset)
}

func (s *EmailService) GetEmail(id string) (*models.Email, error) {
	return s.provider.GetEmail(id)
}

func (s *EmailService) SendEmail(email *models.Email) error {
	return s.provider.SendEmail(email)
}
