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

	"github.com/gorilla/mux"
)

// LoadLoginPage handles the login page request.
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	if utils.IsAuthenticated(r) {
		http.Redirect(w, r, "/feed", http.StatusSeeOther)
		return
	}
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadUsersPage handles the users page request.
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "users.html", nil)
}

// LoadRegisterPage handles the register page request.
func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	if utils.IsAuthenticated(r) {
		http.Redirect(w, r, "/feed", http.StatusSeeOther)
		return
	}
	utils.ExecuteTemplate(w, "register.html", nil)
}

// LoadFeedPage handles the feed page request.
func LoadFeedPage(w http.ResponseWriter, r *http.Request) {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	postsApiURL := fmt.Sprintf("%s/posts", apiURL)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, postsApiURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}
	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	for i := range posts {
		if posts[i].AuthorAvatarURL == "" {
			posts[i].AuthorAvatarURL = "/assets/img/avatar-placeholder.png"
		}
	}

	cookie, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	username, err := utils.GetLoggedUsername(r, apiURL, userID)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to get username"})
		return
	}

	utils.ExecuteTemplate(w, "feed.html", struct {
		Posts    []models.Post
		UserID   uint64
		Username string
	}{
		Posts:    posts,
		UserID:   userID,
		Username: username,
	})
}

// LoadProfilePage retrieves and displays a user's profile page.
func LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	url := fmt.Sprintf("%s/profile/%s", apiURL, username)

	resp, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		responses.HandleStatusCode(w, resp)
		return
	}
	var profileData models.Profile
	if err := json.NewDecoder(resp.Body).Decode(&profileData); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to decode profile data"})
		return
	}

	if profileData.User.AvatarURL == "" {
		profileData.User.AvatarURL = "/assets/img/avatar-placeholder.png"
	}
	if profileData.User.BannerURL == "" {
		profileData.User.BannerURL = "/assets/img/banner-placeholder.png"
	}

	for i := range profileData.Posts {
		if profileData.Posts[i].AuthorAvatarURL == "" {
			profileData.Posts[i].AuthorAvatarURL = "/assets/img/avatar-placeholder.png"
		}
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	username, err = utils.GetLoggedUsername(r, apiURL, userID)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to get username"})
		return
	}
	utils.ExecuteTemplate(w, "profile.html", map[string]interface{}{
		"Profile":  profileData,
		"UserID":   userID,
		"Username": username,
	})
}
