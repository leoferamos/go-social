package routes

import (
	"go_social/internal/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	RequiresAuth: false,
}
