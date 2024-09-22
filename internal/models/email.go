// Package models defines the data structures used in the email client application
package models

import "time"

// Email represents an email message in the system.
type Email struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	Date      time.Time `json:"date"`
	MessageID string    `json:"messageId"`
	ReplyTo   string    `json:"replyTo,omitempty"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
