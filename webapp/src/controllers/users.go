package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreateUser handles user registration.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Invalid user data"})
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Invalid user data"})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	response, err := http.Post(apiURL+"/users", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to create user"})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var loginData struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}
	loginData.Identifier = user.Email
	loginData.Password = user.Password

	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to prepare login data"})
		return
	}

	loginResp, err := http.Post(apiURL+"/login", "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to login after registration"})
		return
	}
	defer loginResp.Body.Close()
	if loginResp.StatusCode >= 400 {
		responses.HandleStatusCode(w, loginResp)
		return
	}

	var authData struct {
		ID    string `json:"id"`
		Token string `json:"token"`
	}
	if err := json.NewDecoder(loginResp.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to decode login response"})
		return
	}

	if err := cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to save auth cookie"})
		return
	}

	http.Redirect(w, r, "/feed", http.StatusSeeOther)
}

// UpdateUser handles user profile updates.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		responses.JSON(w, http.StatusUnauthorized, responses.ErrorAPI{Error: "Not authenticated"})
		return
	}
	userID := cookie["id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Invalid user data"})
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Invalid user data"})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	updateURL := apiURL + "/users/" + userID

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPut, updateURL, bytes.NewBuffer(userJSON))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to update user"})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"message": "User updated successfully"})
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	getURL := apiURL + "/users/" + userID

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, getURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to get user"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var user models.User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to decode user data"})
		return
	}

	if user.AvatarURL == "" {
		user.AvatarURL = "/assets/img/avatar-placeholder.png"
	}
	if user.BannerURL == "" {
		user.BannerURL = "/assets/img/banner-placeholder.png"
	}

	responses.JSON(w, http.StatusOK, user)
}

// FollowUser calls the API to follow a user.
func FollowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	followURL := apiURL + "/users/" + userID + "/follow"
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, followURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to follow user"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]bool{"following": true})
}

// UnfollowUser calls the API to unfollow a user.
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	unfollowURL := apiURL + "/users/" + userID + "/unfollow"
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, unfollowURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to unfollow user"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]bool{"following": false})
}

// GetFollowers retrieves the followers of a user.
func GetFollowers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	getFollowersURL := apiURL + "/users/" + userID + "/followers"

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, getFollowersURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to get followers"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var followers []models.PublicUser
	if err := json.NewDecoder(response.Body).Decode(&followers); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to decode followers"})
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// GetFollowing retrieves the users that a user is following.
func GetFollowing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	getFollowingURL := apiURL + "/users/" + userID + "/following"

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, getFollowingURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to get following"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var following []models.PublicUser
	if err := json.NewDecoder(response.Body).Decode(&following); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to decode following"})
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// IsFollowing checks if the authenticated user is following the specified user.
func IsFollowing(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		responses.JSON(w, http.StatusUnauthorized, responses.ErrorAPI{Error: "Not authenticated"})
		return
	}
	authUserID := cookie["id"]

	vars := mux.Vars(r)
	targetUserID := vars["userId"]

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	isFollowingURL := apiURL + "/users/" + authUserID + "/isFollowing/" + targetUserID

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, isFollowingURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to check following status"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var isFollowing bool
	if err := json.NewDecoder(response.Body).Decode(&isFollowing); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to decode following status"})
		return
	}

	responses.JSON(w, http.StatusOK, map[string]bool{"is_following": isFollowing})
}
