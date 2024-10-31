package routes

import (
	"DevBookAPI/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		URI:                   "/users/{userId}/follow", // id do usuário a ser seguido
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnfollowUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/followers",
		Method:                http.MethodGet,
		Function:              controllers.FindFollowers,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindAllUsers,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.FindOneUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequestAuthentication: true,
	},
}
