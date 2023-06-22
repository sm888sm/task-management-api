package models

// Token represents the authentication token model
type Token struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
