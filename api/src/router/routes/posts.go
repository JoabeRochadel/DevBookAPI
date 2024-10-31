package routes

import (
	"DevBookAPI/src/controllers"
	"net/http"
)

var routePost = []Route{

	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodGet,
		Function:              controllers.FindOnePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodGet,
		Function:              controllers.FindAllPosts,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodPost,
		Function:              controllers.CreatePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePost,
		RequestAuthentication: true,
	},
}
