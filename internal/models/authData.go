package models

// AuthData represents the authentication data for a user
type AuthData struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}
