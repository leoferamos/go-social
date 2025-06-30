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
	{
		URI:          "/register",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/profile/{username}",
		Method:       http.MethodGet,
		Function:     controllers.LoadProfilePage,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{userId}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		RequiresAuth: true,
	},
}
