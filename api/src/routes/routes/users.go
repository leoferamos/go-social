package routes

import "net/http"

var userRoutes = []Route{
	{
		URI:    "users/",
		Method: http.MethodPost,
		Function: func(http.ResponseWriter, *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "users/{userId}",
		Method: http.MethodGet,
		Function: func(http.ResponseWriter, *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "users/",
		Method: http.MethodGet,
		Function: func(http.ResponseWriter, *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "users/{userId}",
		Method: http.MethodPut,
		Function: func(http.ResponseWriter, *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "users/{userId}",
		Method: http.MethodPost,
		Function: func(http.ResponseWriter, *http.Request) {

		},
		RequiresAuth: false,
	},
}
