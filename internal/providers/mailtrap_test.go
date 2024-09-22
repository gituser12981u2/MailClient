package providers

import (
	"os"
	"strconv"
	"testing"

	"mailclient/internal/models"
)

func TestMailtrapSMTPProvider(t *testing.T) {
	// Fetch Mailtrap credentials from environment variables
	host := os.Getenv("MAILTRAP_HOST")
	portStr := os.Getenv("MAILTRAP_PORT")
	username := os.Getenv("MAILTRAP_USERNAME")
	password := os.Getenv("MAILTRAP_PASSWORD")

	if host == "" || username == "" || password == "" {
		t.Fatalf("Mailtrap credentials not set in environment variables")
	}

	// Create an SMTP provider with Mailtrap credentials
	config := models.SMTPConfig{
		Host:     host,
		Username: username,
		Password: password,
	}

	// Convert port string to int
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatalf("Invalid port number: %v", err)
		}
		config.Port = port
	}

	provider, err := NewSMTPProvider(config)
	if err != nil {
		t.Fatalf("Failed to create SMTP provider: %v", err)
	}

	// Test SendEmail
	email := &models.Email{
		Subject: "Test Subject",
		To:      "sender@example.com",
		Body:    "This is a test email sent via Mailtrap with TLS",
	}
	err = provider.SendEmail(email)
	if err != nil {
		t.Errorf("SendEmail failed: %v", err)
	}

	t.Log("Check Mailtrap inbox for the test email")
}
