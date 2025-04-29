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
}
