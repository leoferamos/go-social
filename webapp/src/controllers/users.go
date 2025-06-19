package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CreateUser handles user registration.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user map[string]string
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user data"})
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user data"})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	response, err := http.Post(apiURL+"/users", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	loginData := map[string]string{
		"identifier": user["email"],
		"password":   user["password"],
	}
	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to prepare login data"})
		return
	}

	loginResp, err := http.Post(apiURL+"/login", "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to login after registration"})
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
		responses.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to decode login response"})
		return
	}

	if err := cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to save auth cookie"})
		return
	}

	http.Redirect(w, r, "/feed", http.StatusSeeOther)
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
	cookie, _ := cookies.Read(r)
	userID := cookie["id"]
	utils.ExecuteTemplate(w, "profile.html", map[string]interface{}{
		"Profile": profileData,
		"UserID":  userID,
	})
}
