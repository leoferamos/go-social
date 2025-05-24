package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"webapp/src/responses"
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
	responses.JSON(w, response.StatusCode, nil)
}
