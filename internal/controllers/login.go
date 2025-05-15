package controllers

import (
	"encoding/json"
	"errors"
	"go_social/internal/auth"
	"go_social/internal/db"
	"go_social/internal/models"
	"go_social/internal/repositories"
	"go_social/internal/responses"
	"go_social/internal/security"
	"io"
	"net/http"
)

// Login handles user login requests.
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, errors.New("error connecting to database"))
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userFromDB, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, errors.New("error fetching user"))
		return
	}
	if err = security.CheckPasswordHash(userFromDB.Password, user.Password); err != nil {
		responses.JSONError(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	token, err := auth.CreateToken(int(userFromDB.ID))
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, errors.New("error generating token"))
		return
	}
	w.Write([]byte(token))
}
