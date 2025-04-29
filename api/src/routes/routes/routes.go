package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents an API route.
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// userRoutes contains all user-related routes.
func Configure(r *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
