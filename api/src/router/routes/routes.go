package routes

import (
	"DevBookAPI/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	var routes []Route

	routes = append(routes, routeUsers...)
	routes = append(routes, routeLogin)
	routes = append(routes, routePost...)

	for _, route := range routes {
		if route.RequestAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Function),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	return r
}
