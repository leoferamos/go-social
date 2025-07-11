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
	"strconv"
	"strings"
	"time"

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

	if err = user.Prepare("registration"); err != nil {
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

// GetUser retrieves a user from the database by ID
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
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}

	userIDFromToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	if userID != userIDFromToken {
		responses.JSONError(w, http.StatusForbidden, errors.New("you cannot update another user"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
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
	if err = repository.UpdateUser(userID, user); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}

	userIDFromToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	if userID != userIDFromToken {
		responses.JSONError(w, http.StatusForbidden, errors.New("you cannot delete another user"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.DeleteUser(userID); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser allows a user to follow another user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if followerID == userID {
		responses.JSONError(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	if err = repository.FollowUser(userID, followerID); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser allows a user to unfollow another user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if followerID == userID {
		responses.JSONError(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	if err = repository.UnfollowUser(userID, followerID); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// GetFollowers retrieves the followers of a user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
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
	followers, err := repository.GetFollowers(userID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}

// GetFollowing retrieves the users that a user is following
func GetFollowing(w http.ResponseWriter, r *http.Request) {
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
	following, err := repository.GetFollowing(userID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, following)
}

// IsFollowing checks if the authenticated user is following the specified user
func IsFollowing(w http.ResponseWriter, r *http.Request) {
	userIDFromToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
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
	isFollowing, err := repository.IsFollowing(userID, userIDFromToken)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]bool{"is_following": isFollowing})
}

// ResetPassword resets the password of a user
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	userIDFromToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if userID != userIDFromToken {
		responses.JSONError(w, http.StatusForbidden, errors.New("you cannot reset another user's password"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if err = models.ValidatePasswordStrength(password.NewPassword); err != nil {
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
	passwordOnDB, err := repository.GetPasswordByID(userID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.CheckPasswordHash(passwordOnDB, password.Password); err != nil {
		responses.JSONError(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}
	hashedPassword, err := security.HashPassword(password.NewPassword)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	if err = repository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	username := parameters["username"]

	dbConn, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	repository := repositories.NewUsersRepository(dbConn)
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		responses.JSONError(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// GetProfileByUsername retrieves a user profile (aggregated info) by username
func GetProfileByUsername(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	username := parameters["username"]

	dbConn, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	repository := repositories.NewUsersRepository(dbConn)
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		responses.JSONError(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	postsRepo := repositories.NewPostsRepository(dbConn)
	var posts []models.Post
	var postsErr error

	userIDFromToken, err := auth.ExtractUserID(r)

	isFollowing := false
	if err == nil && userIDFromToken != user.ID {
		isFollowing, err = repository.IsFollowing(user.ID, userIDFromToken)
		if err != nil {
			responses.JSONError(w, http.StatusInternalServerError, err)
			return
		}
	}

	posts, postsErr = postsRepo.GetUserPosts(user.ID, userIDFromToken)
	if postsErr != nil {
		responses.JSONError(w, http.StatusInternalServerError, postsErr)
		return
	}

	followers, err := repository.GetFollowers(user.ID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	following, err := repository.GetFollowing(user.ID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	publicUser := models.PublicUser{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		Followers: len(followers),
		Following: len(following),
	}

	profile := models.Profile{
		User:        publicUser,
		IsFollowing: isFollowing,
		Posts:       posts,
	}
	responses.JSON(w, http.StatusOK, profile)
}
