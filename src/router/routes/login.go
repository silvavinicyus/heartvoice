package routes

import (
	"heartvoice/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		Uri:                    "/sessions",
		Method:                 http.MethodPost,
		Function:               controllers.Login,
		RequiresAuthentication: false,
	},
}
