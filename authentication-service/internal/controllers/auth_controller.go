package controllers

import (
	"context"

	"github.com/sm888sm/task-management-api/authentication-service/internal/models"
	"github.com/sm888sm/task-management-api/authentication-service/internal/proto/auth"
	"github.com/sm888sm/task-management-api/authentication-service/internal/services"
)

// AuthController represents the authentication controller

type AuthController struct {
	authService  *services.AuthService
	tokenService *services.TokenService
}

// NewAuthController creates a new instance of the AuthController
func NewAuthController(authService *services.AuthService, tokenService *services.TokenService) *AuthController {
	return &AuthController{
		authService:  authService,
		tokenService: tokenService,
	}
}

// Login handles the login request
func (c *AuthController) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := c.authService.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := c.tokenService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		User:  &auth.User{Id: user.ID, Username: user.Username},
		Token: token,
	}, nil
}

// Register handles the user registration request
func (c *AuthController) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	err := c.authService.RegisterUser(user.Username, user.Password, user.Email)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}

// ValidateToken handles the token validation request
func (c *AuthController) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	tokenString := req.Token

	claims, err := c.tokenService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return &auth.ValidateTokenResponse{
		UserId: claims["userID"].(int32),
	}, nil
}
