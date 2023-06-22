package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sm888sm/task-management-api/authentication-service/config"
	"github.com/sm888sm/task-management-api/authentication-service/internal/utils"
)

// AuthenticationMiddleware is a middleware for authentication
type AuthenticationMiddleware struct {
	config *config.Config
}

// NewAuthenticationMiddleware creates a new instance of the AuthenticationMiddleware
func NewAuthenticationMiddleware(config *config.Config) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		config: config,
	}
}

// Middleware handles the authentication middleware logic
func (am *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractTokenFromRequest(r)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := validateToken(token, am.config.JWTSecretKey)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// extractTokenFromRequest extracts the JWT token from the request's Authorization header
func extractTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ""
	}

	return tokenParts[1]
}

// validateToken validates the JWT token and returns the user ID if valid
func validateToken(token, secretKey string) (int, error) {
	claims, err := utils.VerifyJWTToken(token, secretKey)
	if err != nil {
		return 0, err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, err
	}

	return int(userID), nil
}
