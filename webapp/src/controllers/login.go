package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginPage handles the login page request.
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
