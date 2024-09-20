package providers

import (
	"fmt"

	"mailclient/internal/models"
)

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

func New(providerType string, config map[string]string) (EmailProvider, error) {
	switch providerType {
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
