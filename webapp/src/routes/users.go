package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:          "/register",
		Method:       http.MethodGet,
		Function:     controllers.LoadRegisterPage,
		RequiresAuth: false,
	},
}
