package routes

import (
	"github.com/gorilla/mux"
)

// Generate returns a router with all application routes registered.
func Generate() *mux.Router {
	r := mux.NewRouter()
	return Configure(r)
}
