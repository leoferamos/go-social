package utils

import (
	"net/http"
	"webapp/src/cookies"
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
