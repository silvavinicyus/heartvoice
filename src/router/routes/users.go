package routes

import (
	"heartvoice/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:                    "/auth/register",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUsers,
		RequiresAuthentication: false,
	},
}
