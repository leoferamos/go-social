package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routes = append(routes, feedRoute...)

	for _, route := range routes {

		if route.RequiresAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.AuthMiddleware(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
