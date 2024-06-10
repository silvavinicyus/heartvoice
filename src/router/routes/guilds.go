package routes

import (
	"heartvoice/src/controllers"
	"net/http"
)

var guildRoutes = []Route{
	{
		Uri:                    "/guilds",
		Method:                 http.MethodPost,
		Function:               controllers.CreateGuild,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/guilds",
		Method:                 http.MethodGet,
		Function:               controllers.FindAllGuilds,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/guilds/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindByGuild,
		RequiresAuthentication: true,
	},
}
