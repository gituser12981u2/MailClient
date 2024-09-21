// Package models defines the data structures used in the email client application
package models

// Email represents an email message in the system.
type Email struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Sender  string `json:"sender"`
	Body    string `json:"body"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
