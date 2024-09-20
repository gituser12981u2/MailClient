package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerAddress  string            `json:"server_address"`
	Provider       string            `json:"provider"`
	ProviderConfig map[string]string `json:"provider_config"`
	JWTSecret      string            `json:"jwt_secret"`
}

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
