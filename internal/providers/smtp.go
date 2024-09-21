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

func (p *SMTPProvider) SendEmail(email *models.Email) error {
	log.Println("SMTP Provider: SendEmail called")

	err := p.connectAndSend(email)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (p *SMTPProvider) connectAndSend(email *models.Email) error {
	ports := []int{587, 465, 2525, 25}

	if p.config.Port != 0 {
		ports = append([]int{p.config.Port}, ports...)
		// Remove duplicate port if user specified one of the default ports
		for i, port := range ports[1:] {
			if port == p.config.Port {
				ports = append(ports[:i+1], ports[i+2:]...)
				break
			}
		}
	}

	var lastErr error
	for _, port := range ports {
		err := p.tryConnect(port, email)
		if err == nil {
			return nil
		}
		lastErr = err
	}

	return fmt.Errorf("failed to connect and send email on default ports: %v", lastErr)
}

func (p *SMTPProvider) tryConnect(port int, email *models.Email) error {
	addr := fmt.Sprintf("%s:%d", p.config.Host, port)

	var c *smtp.Client
	var err error

	if port == 465 {
		// Implicit TLS
		conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: p.config.Host})
		if err != nil {
			return fmt.Errorf("failed to establish TLS connection on port 465: %w", err)
		}
		c, err = smtp.NewClient(conn, p.config.Host)
	} else {
		// Start with plain connection
		c, err = smtp.Dial(addr)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server on port %d: %w", port, err)
	}
	defer c.Close()

	if port != 465 {
		if ok, _ := c.Extension("STARTTLS"); ok {
			config := &tls.Config{ServerName: p.config.Host}
			if err = c.StartTLS(config); err != nil {
				return fmt.Errorf("failed to start TLS on port %d: %w", port, err)
			}
		}
	}
	if err = c.Auth(p.auth); err != nil {
		return fmt.Errorf("authentication failed on port %d: %w", port, err)
	}

	return p.sendEmail(c, email)
}

func (p *SMTPProvider) sendEmail(c *smtp.Client, email *models.Email) error {
	if err := c.Mail(p.config.Username); err != nil {
		return fmt.Errorf("MAIL command failed: %w", err)
	}

	to := strings.Split(email.Sender, ",")
	for _, recipient := range to {
		if err := c.Rcpt(strings.TrimSpace(recipient)); err != nil {
			return fmt.Errorf("RCPT command failed: %w", err)
		}
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("DATA command failed: %w", err)
	}

	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email.Sender, email.Subject, email.Body)
	if _, err = w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write email body: %w", err)
	}

	if err = w.Close(); err != nil {
		return fmt.Errorf("failed to close email body writer: %w", err)
	}

	return c.Quit()
}

func (p *SMTPProvider) Connect() error {
	// No prior persistent connection needed
	return nil
}

func (p *SMTPProvider) Disconnect() error {
	// No persistent connection to close
	return nil
}

func (p *SMTPProvider) GetEmails(folder string, limit, offset int) ([]models.Email, error) {
	return nil, fmt.Errorf("GetEmails not implemented for SMTP provider. Use IMAP provider for this functionality")
}

func (p *SMTPProvider) GetEmail(id string) (*models.Email, error) {
	return nil, fmt.Errorf("GetEmail not supported by SMTP provider. Use IMAP provider for this functionality")
}

func (p *SMTPProvider) DeleteEmail(id string) error {
	return fmt.Errorf("DeleteEmail not supported by SMTP. Use IMAP provider for this functionality")
}
