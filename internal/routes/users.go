package routes

import (
	"go_social/internal/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
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
		URI:          "/users/{userId}/reset-password",
		Method:       http.MethodPost,
		Function:     controllers.ResetPassword,
		RequiresAuth: true,
	},
	{
		URI:          "/profile/{username}",
		Method:       http.MethodGet,
		Function:     controllers.GetUserByUsername,
		RequiresAuth: false,
	},
}
