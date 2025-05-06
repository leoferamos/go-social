package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user in the system.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare calls the methods to format and validate the user data.
func (u *User) Prepare(context string) error {
	if err := u.validate(context); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validate(context string) error {
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if context == "registration" && u.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
}
