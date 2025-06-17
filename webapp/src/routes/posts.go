package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRoute = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{id}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{id}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePost,
		RequiresAuth: true,
	},
}
