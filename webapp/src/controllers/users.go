package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
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
	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(userJSON))
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
