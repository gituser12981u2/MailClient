// Package config provides functionality for loading and managing application configuration.
package config

import (
	"encoding/json"
	"os"
)

// ProviderConfig represents the configuration for an email provider
type ProviderConfig struct {
	Type   string            `json:"type"`
	Config map[string]string `json:"config"`
}

// Config represents the application configuration
type Config struct {
	ServerAddress  string         `json:"server_address"`
	JWTSecret      string         `json:"jwt_secret"`
	ProviderConfig ProviderConfig `json:"provider_config"`
}

// Load reads the configuration from a JSON file and returns a Config struct.
// It returns an error if the file cannot be read or parsed.
func Load() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
