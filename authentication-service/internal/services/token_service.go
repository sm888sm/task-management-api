package services

import (
	"fmt"
	"time"

	"github.com/sm888sm/task-management-api/authentication-service/config"
	"github.com/sm888sm/task-management-api/authentication-service/internal/models"
	"github.com/sm888sm/task-management-api/authentication-service/internal/utils"
)

// TokenService represents the token service
type TokenService struct {
	config *config.Config
}

// NewTokenService creates a new instance of the TokenService
func NewTokenService(config *config.Config) *TokenService {
	return &TokenService{
		config: config,
	}
}

// GenerateToken generates a new authentication token for the given user
func (ts *TokenService) GenerateToken(user *models.User) (string, error) {
	// Set the token expiry time
	expirationTime := time.Now().Add(ts.config.TokenExpirationTime)

	// Create the claims for the token
	claims := map[string]interface{}{
		"userID": user.ID,
		"exp":    expirationTime.Unix(),
	}

	// Generate the JWT token
	token, err := utils.GenerateJWTToken(claims, ts.config.JWTSecretKey, ts.config.TokenExpirationTime)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (ts *TokenService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	// TODO Implement this method
	claims, err := utils.VerifyJWTToken(tokenString, ts.config.JWTSecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return claims, nil
}
