package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// Login handles user login.
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"identifier": r.FormValue("identifier"),
		"password":   r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}

	response, err := http.Post(apiURL+"/login", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var authData models.AuthData
	if err := json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	if err := cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	responses.JSON(w, http.StatusOK, authData)

}
