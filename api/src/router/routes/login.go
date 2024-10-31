package routes

import (
	"DevBookAPI/src/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequestAuthentication: false,
}
