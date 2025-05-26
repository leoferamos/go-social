package routes

import "webapp/src/controllers"

var feedRoute = []Route{
	{
		URI:          "/feed",
		Method:       "GET",
		Function:     controllers.LoadFeedPage,
		RequiresAuth: true,
	},
}
