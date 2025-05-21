package controllers

import "net/http"

// LoadLoginPage handles the login page request.
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}
