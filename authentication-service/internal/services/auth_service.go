package services

import (
	"github.com/sm888sm/task-management-api/authentication-service/internal/models"
	"github.com/sm888sm/task-management-api/authentication-service/internal/repositories"
	"github.com/sm888sm/task-management-api/authentication-service/internal/utils"
)

// AuthService represents the authentication service
type AuthService struct {
	userRepository repositories.UserRepository
}

// NewAuthService creates a new instance of the AuthService
func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

// RegisterUser registers a new user with the provided details
func (as *AuthService) RegisterUser(username, password, email string) error {
	// Check if the user already exists
	_, err := as.userRepository.GetUserByUsername(username)
	if err == nil {
		return utils.NewError("username already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return utils.NewError("failed to hash password")
	}

	// Create a new user
	user := models.NewUser(username, hashedPassword, email)

	// Save the user in the database
	err = as.userRepository.CreateUser(user)
	if err != nil {
		return utils.NewError("failed to create user")
	}

	return nil
}

// AuthenticateUser authenticates a user with the provided credentials
func (as *AuthService) AuthenticateUser(username, password string) (*models.User, error) {
	// Get the user from the database
	user, err := as.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, utils.NewError("user not found")
	}

	// Verify the password
	isValid := utils.VerifyPassword(password, user.Password)
	if !isValid {
		return nil, utils.NewError("incorrect password")
	}

	return user, nil
}
