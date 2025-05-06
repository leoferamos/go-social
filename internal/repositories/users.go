package repositories

import (
	"database/sql"
	"fmt"
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
func (repository *users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)")
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

// GetUsers retrieves users from the database based on the provided name or username
func (repository *users) GetUsers(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	rows, err := repository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE name LIKE ? OR username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
