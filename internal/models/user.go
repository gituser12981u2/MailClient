// Package models defines the data structures used in the email client application.
package models

// user represents a user in the system
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
