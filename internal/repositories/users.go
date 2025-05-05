package repositories

import (
	"database/sql"
	"go_social/internal/models"
)

type users struct {
	db *sql.DB
}

// NewUsersRepository creates a new instance of the users repository
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// CreateUser creates a new user in the database
func (u *users) CreateUser(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), nil
}
