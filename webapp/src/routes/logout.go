package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRoute = []Route{
	{
		URI:          "/logout",
		Method:       http.MethodGet,
		Function:     controllers.Logout,
		RequiresAuth: true,
	},
}
