package utils

import "fmt"

// AppError represents an application error
type AppError struct {
	Message string
}

// NewError creates a new application error with the provided message
func NewError(message string) *AppError {
	return &AppError{
		Message: message,
	}
}

// Error returns the error message
func (e *AppError) Error() string {
	return e.Message
}

// WrapError wraps an error with a custom message
func WrapError(err error, message string) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%s: %v", message, err),
	}
}
