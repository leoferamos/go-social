package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

// LoadLoginPage handles the login page request.
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadUsersPage handles the users page request.
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "users.html", nil)
}

// LoadRegisterPage handles the register page request.
func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

// LoadFeedPage handles the feed page request.
func LoadFeedPage(w http.ResponseWriter, r *http.Request) {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	apiURL = fmt.Sprintf("%s/posts", apiURL)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, apiURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}
	var posts []models.Posts
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "feed.html", struct {
		Posts  []models.Posts
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}
