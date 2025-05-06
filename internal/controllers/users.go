package controllers

import (
	"encoding/json"
	"go_social/internal/db"
	"go_social/internal/models"
	"go_social/internal/repositories"
	"go_social/internal/responses"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

// GetUsers retrieves users from the database based on the provided name or username
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, err := repository.GetUsers(nameOrUsername)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// GetUsers retrieves all users from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
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
	user, err := repository.GetUserByID(userID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates an existing user in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
