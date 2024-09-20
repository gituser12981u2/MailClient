// Package config provides functionality for loading and managing application configuration.
package config

import (
	"encoding/json"
	"os"
)

// Config represents the application configuration
type Config struct {
	ServerAddress  string            `json:"server_address"`
	Provider       string            `json:"provider"`
	ProviderConfig map[string]string `json:"provider_config"`
	JWTSecret      string            `json:"jwt_secret"`
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
