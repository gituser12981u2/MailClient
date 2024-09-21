// Package providers defines the EmailProvider interface and factory function.
package providers

import (
	"fmt"
	"strconv"

	"mailclient/internal/models"
)

// EmailProvider defines the interface for email provider implementations.
type EmailProvider interface {
	Connect() error
	Disconnect() error
	GetEmails(folder string, limit, offset int) ([]models.Email, error)
	GetEmail(id string) (*models.Email, error)
	SendEmail(email *models.Email) error
	DeleteEmail(id string) error
	// MarkAsRead(id string) error
	// MarkAsUnread(id string) error
	// GetFolders() ([]string, error)
	// MoveEmail(id, folder string) error
	// SearchEmails(query string, limit, offset int) ([]models.Email, error)
}

// New creates a new EmailProvider based on the provided type and configuration.
func New(providerType string, config map[string]string) (EmailProvider, error) {
	switch providerType {
	case "smtp":
		smtpConfg := models.SMTPConfig{
			Host:     config["host"],
			Username: config["username"],
			Password: config["password"],
		}
		if port, err := strconv.Atoi(config["port"]); err == nil {
			smtpConfg.Port = port
		}
		return NewSMTPProvider(smtpConfg)
	case "proton":
		provider, err := NewProtonMailProvider(config["username"], config["password"])
		if err != nil {
			return nil, fmt.Errorf("failed to create ProtonMail provider: %w", err)
		}
		return provider, nil
	default:
		return nil, fmt.Errorf("unsupported email provider: %s", providerType)
	}
}
