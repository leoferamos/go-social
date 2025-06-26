package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/cookies"
	"webapp/src/requests"
)

// IsAuthenticated checks if the request has a valid auth_data cookie with a non-empty token.
func IsAuthenticated(r *http.Request) bool {
	cookie, err := cookies.Read(r)
	if err != nil {
		return false
	}
	token, ok := cookie["token"]
	return ok && token != ""
}

func GetLoggedUsername(r *http.Request, apiURL string, userID uint64) (string, error) {
	userURL := fmt.Sprintf("%s/users/%d", apiURL, userID)
	userResp, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, userURL, nil)
	if err != nil || userResp.StatusCode >= 400 {
		return "", err
	}
	defer userResp.Body.Close()
	var userData struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(userResp.Body).Decode(&userData); err != nil {
		return "", err
	}
	return userData.Username, nil
}
