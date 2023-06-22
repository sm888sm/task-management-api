package models

import (
	"time"
)

// User represents the user model
type User struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new User instance with the given details
func NewUser(username, password, email string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
