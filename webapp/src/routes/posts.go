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
}
