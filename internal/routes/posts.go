package routes

import (
	"go_social/internal/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetPost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Function:     controllers.GetPosts,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		RequiresAuth: true,
	},
}
