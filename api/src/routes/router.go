package routes

import (
	"api/src/routes/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with all application routes registered.
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
