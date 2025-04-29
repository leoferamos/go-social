package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "users/",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		RequiresAuth: false,
	},
	{
		URI:          "users/",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		RequiresAuth: false,
	},
	{
		URI:          "users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI:          "users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiresAuth: false,
	},
}
