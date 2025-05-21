package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents an web application route.
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// Configure sets up the routes for the web application.
func Configure(r *mux.Router) *mux.Router {
	routes := loginRoute
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)

	}

	fileServer := http.FileServer(http.Dir("./assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
