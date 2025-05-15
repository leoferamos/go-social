package models

// Password represents a password model.
type Password struct {
	NewPassword string `json:"new_password"`
	Password    string `json:"password"`
}
