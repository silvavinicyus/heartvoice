package routes

import (
	"heartvoice/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		Uri:                    "/auth/login",
		Method:                 http.MethodPost,
		Function:               controllers.Login,
		RequiresAuthentication: false,
	},
}
