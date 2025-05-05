package controllers

import (
	"encoding/json"
	"fmt"
	"go_social/internal/db"
	"go_social/internal/models"
	"go_social/internal/repositories"
	"io"
	"log"
	"net/http"
)

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var usuario models.User
	if err = json.Unmarshal(bodyRequest, &usuario); err != nil {
		log.Fatal(err)
	}
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepository(db)
	userId, err := repository.CreateUser(usuario)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("User created successfully with ID: %d", userId)))
}

// GetUser retrieves a user from the database by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
}

// GetUsers retrieves all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of all users"))
}

// UpdateUser updates an existing user in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
