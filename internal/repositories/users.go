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

// GetUserByID retrieves a user from the database by ID
func (repository *users) GetUserByID(userID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func (repository *users) UpdateUser(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database
func (repository *users) DeleteUser(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user from the database by email and return the id and hashed password
func (repository *users) GetUserByEmail(email string) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// FollowUser allows a user to follow another user
func (repository *users) FollowUser(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

// UnfollowUser allows a user to unfollow another user
func (repository *users) UnfollowUser(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

// GetFollowers retrieves the followers of a user
func (repository *users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.created_at
		FROM users u inner join followers f on u.id = f.follower_id where f.user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}

	return followers, nil
}

// GetFollowing retrieves the users that a user is following
func (repository *users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.created_at
		FROM users u inner join followers f on u.id = f.user_id where f.follower_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		following = append(following, user)
	}

	return following, nil
}

// GetPasswordByID retrieves a user's password by ID
func (repository *users) GetPasswordByID(userID uint64) (string, error) {
	rows, err := repository.db.Query(
		"SELECT password FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.Password); err != nil {
			return "", err
		}
	}
	return user.Password, nil
}

// UpdatePassword updates a user's password in the database
func (repository *users) UpdatePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(password, userID); err != nil {
		return err
	}
	return nil
}
