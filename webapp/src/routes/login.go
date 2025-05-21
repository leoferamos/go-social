package routes

import (
	"net/http"

	"webapp/src/controllers"
)

var loginRoute = []Route{
	{
		URI:          "/",
		Method:       http.MethodPost,
		Function:     controllers.LoadLoginPage,
		RequiresAuth: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.LoadLoginPage,
		RequiresAuth: false,
	},
}
