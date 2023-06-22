package config

import (
	"fmt"
	"os"
	"time"
)

// Config represents the configuration parameters for the authentication service
type Config struct {
	Port                string
	DatabaseURL         string
	JWTSecretKey        string
	PasswordSalt        string
	TokenExpirationTime time.Duration
}

// LoadConfig loads the configuration parameters from environment variables
func LoadConfig() (*Config, error) {
	port := os.Getenv("AUTH_SERVICE_PORT")
	if port == "" {
		port = "50051" // Default port if not specified
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY is not set")
	}

	passwordSalt := os.Getenv("PASSWORD_SALT")
	if passwordSalt == "" {
		return nil, fmt.Errorf("PASSWORD_SALT is not set")
	}

	config := &Config{
		Port:         port,
		DatabaseURL:  databaseURL,
		JWTSecretKey: jwtSecretKey,
		PasswordSalt: passwordSalt,
	}

	return config, nil
}
