// Package providers implements different email provider integrations
package providers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"

	"mailclient/internal/models"

	"github.com/ProtonMail/go-proton-api"
)

// ProtonMailProvider implements the EmailProvider interface for ProtonMail.
type ProtonMailProvider struct {
	username string
	password string
	manager  *proton.Manager
	client   *proton.Client
	mu       sync.Mutex
}

// NewProtonMailProvider creates a new ProtonMailProvider instance.
func NewProtonMailProvider(username, password string) (*ProtonMailProvider, error) {
	provider := &ProtonMailProvider{
		username: username,
		password: password,
		manager:  proton.New(proton.WithHostURL("https://api.protonmail.ch"), proton.WithAppVersion("Other")),
	}

	if err := provider.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	return provider, nil
}

// Connect establishes a connection to the ProtonMail API.
func (p *ProtonMailProvider) Connect() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.client != nil {
		return nil // Already connected
	}

	ctx := context.Background()

	// Login
	client, auth, err := p.manager.NewClientWithLogin(ctx, p.username, []byte(p.password))
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	// Handle 2FA if enabled
	if auth.TwoFA.Enabled&proton.HasTOTP != 0 {
		fmt.Print("Enter 2FA code: ")
		reader := bufio.NewReader(os.Stdin)
		twoFACode, _ := reader.ReadString('\n')
		twoFACode = twoFACode[:len(twoFACode)-1] // Remove newline character

		if err := client.Auth2FA(ctx, proton.Auth2FAReq{TwoFactorCode: twoFACode}); err != nil {
			return fmt.Errorf("2FA authentication failed: %v", err)
		}
	}

	p.client = client
	return nil
}

// Disconnect closes the connection to the ProtonMail API.
func (p *ProtonMailProvider) Disconnect() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.client != nil {
		// TODO: Perform cleanup here
		p.client = nil
	}
	return nil
}

// GetEmails retrieves emails from the specified folder with pagination.
func (p *ProtonMailProvider) GetEmails(folder string, limit, offset int) ([]models.Email, error) {
	ctx := context.Background()
	filter := proton.MessageFilter{}
	metadata, err := p.client.GetMessageMetadata(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metadata: %v", err)
	}

	// Convert messages to Email structure
	emails := make([]models.Email, len(metadata))
	for i, msg := range metadata {
		emails[i] = models.Email{
			ID:      msg.ID,
			Subject: msg.Subject,
			To:      msg.Sender.Address,
		}
	}

	return emails, nil
}

// GetEmail retrieves a single email by its ID.
func (p *ProtonMailProvider) GetEmail(id string) (*models.Email, error) {
	if err := p.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	// TODO: Implement this method
	return nil, fmt.Errorf("GetEmail not implemented")
}

// SendEmail sends a new email.
func (p *ProtonMailProvider) SendEmail(email *models.Email) error {
	if err := p.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	// TODO: Implement this method
	return fmt.Errorf("SendEmail not implemented")
}

// DeleteEmail deletes an email by its ID.
func (p *ProtonMailProvider) DeleteEmail(id string) error {
	if err := p.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	// TODO: Implement this method
	return fmt.Errorf("DeleteEmail not implemented")
}
