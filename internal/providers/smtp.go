package providers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"mailclient/internal/models"
)

type SMTPProvider struct {
	config models.SMTPConfig
	auth   smtp.Auth
}

func NewSMTPProvider(config models.SMTPConfig) (*SMTPProvider, error) {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	return &SMTPProvider{
		config: config,
		auth:   auth,
	}, nil
}

func (p *SMTPProvider) Connect() error {
	ports := []int{587, 465, 2525, 25} // Default port order

	if p.config.Port != 0 {
		ports = append([]int{p.config.Port}, ports...)
	}

	var lastErr error
	for _, port := range ports {
		addr := fmt.Sprintf("%s:%d", p.config.Host, port)

		var c *smtp.Client
		var err error

		if port == 465 {
			// Implicit TLS
			conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: p.config.Host})
			if err != nil {
				lastErr = err
				continue
			}
			c, err = smtp.NewClient(conn, p.config.Host)
		} else {
			// Start with plain connection
			c, err = smtp.Dial(addr)
		}

		if err != nil {
			lastErr = err
			continue
		}
		defer c.Close()

		if port != 465 {
			if ok, _ := c.Extension("STARTTLS"); ok {
				config := &tls.Config{ServerName: p.config.Host}
				if err = c.StartTLS(config); err != nil {
					return fmt.Errorf("failed to start TLS: %w", err)
				}
			}
		}
		if err = c.Auth(p.auth); err != nil {
			lastErr = err
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to connect to SMTP server on any port: %v", lastErr)
}

func (p *SMTPProvider) Disconnect() error {
	// No persistent connection to close
	return nil
}

func (p *SMTPProvider) GetEmails(folder string, limit, offset int) ([]models.Email, error) {
	return nil, fmt.Errorf("GetEmails not implemented for SMTP provider")
}

func (p *SMTPProvider) GetEmail(id string) (*models.Email, error) {
	return nil, fmt.Errorf("GetEmail not implemented for SMTP provider")
}

func (p *SMTPProvider) SendEmail(email *models.Email) error {
	log.Println("SMTP Provider: SendEmail called")
	addr := fmt.Sprintf("%s:%d", p.config.Host, p.config.Port)

	c, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: p.config.Host}
		if err = c.StartTLS(config); err != nil {
			return fmt.Errorf("failed to start TLS: %w", err)
		}
	}

	if err = c.Auth(p.auth); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	if err = c.Mail(p.config.Username); err != nil {
		return fmt.Errorf("MAIL command failed: %w", err)
	}

	to := strings.Split(email.Sender, ",")
	for _, recipient := range to {
		if err = c.Rcpt(strings.TrimSpace(recipient)); err != nil {
			return fmt.Errorf("RCPT command failed: %w", err)
		}
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("DATA command failed: %w", err)
	}

	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email.Sender, email.Subject, email.Body)
	_, err = w.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("failed to write email body: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close email body writer: %w", err)
	}

	return c.Quit()
}

func (p *SMTPProvider) DeleteEmail(id string) error {
	return fmt.Errorf("DeleteEmail not implemented for SMTP provider")
}
