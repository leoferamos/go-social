package models

type LoginRequest struct {
    Identifier string `json:"identifier"`
    Password   string `json:"password"`
}