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
	"strings"
)

// Login handles user login requests.
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var loginReq models.LoginRequest
	if err = json.Unmarshal(bodyRequest, &loginReq); err != nil {
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

	var userFromDB models.User
	if strings.Contains(loginReq.Identifier, "@") {
		userFromDB, err = repository.GetUserByEmail(loginReq.Identifier)
	} else {
		userFromDB, err = repository.GetUserByUsername(loginReq.Identifier)
	}
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}
	if err = security.CheckPasswordHash(userFromDB.Password, loginReq.Password); err != nil {
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
