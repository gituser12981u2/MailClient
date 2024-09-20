// Package services provides business logic for email client application.
package services

// AuthService handles authentication-related operations.
type AuthService struct {
	jwtSecret []byte
}

// NewAuthService creates a new AuthService instance.
func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{
		jwtSecret: []byte(jwtSecret),
	}
}
