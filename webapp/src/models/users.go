package models

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Bio      string `json:"bio,omitempty"`
}
