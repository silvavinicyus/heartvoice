package routes

import (
	"heartvoice/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
