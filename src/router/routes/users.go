package routes

import (
	"heartvoice/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUsers,
		RequiresAuthentication: false,
	},
}
