package repositories

import (
	"database/sql"
	"fmt"

	"github.com/sm888sm/task-management-api/authentication-service/internal/models"
)

// UserRepository represents the repository for user operations
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of the UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetUserByUsername retrieves a user from the database by username
func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, email FROM users WHERE username = ?"
	row := ur.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	_, err := ur.db.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}
