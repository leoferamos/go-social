package models

import (
	"errors"
	"go_social/internal/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user in the system.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare calls the methods to format and validate the user data.
func (u *User) Prepare(context string) error {
	if err := u.validate(context); err != nil {
		return err
	}
	if err := u.format(context); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(context string) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email format")
	}
	if context == "registration" && u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (u *User) format(context string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	if context == "registration" {
		hashedPassword, err := security.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}
	return nil
}
