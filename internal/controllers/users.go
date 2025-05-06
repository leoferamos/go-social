package controllers

import (
	"encoding/json"
	"go_social/internal/db"
	"go_social/internal/models"
	"go_social/internal/repositories"
	"go_social/internal/responses"
	"io"
	"net/http"
)

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userId, err := repository.CreateUser(user)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, userId)
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
