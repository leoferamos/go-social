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
		URI:          "/profile/{username}",
		Method:       http.MethodGet,
		Function:     controllers.LoadProfilePage,
		RequiresAuth: false,
	},
	{
		URI:          "/register",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUserByID,
		RequiresAuth: true,
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
	{
		URI:          "/users/{userId}/followers",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/following",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowing,
		RequiresAuth: true,
	},
	{
		URI:          "/users/isFollowing/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.IsFollowing,
		RequiresAuth: true,
	},
}
