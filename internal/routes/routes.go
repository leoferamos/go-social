package routes

import (
	"go_social/internal/middlewares"
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

// Configure sets up the routes for the application.
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequiresAuth {
			r.HandleFunc(route.URI, middlewares.AuthMiddleware(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}
	}
	return r
}
