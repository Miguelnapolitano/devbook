package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.ListUsers,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Function: controllers.RetrieveUser,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}/follow",
		Method: http.MethodPost,
		Function: controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}/unfollow",
		Method: http.MethodPost,
		Function: controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}/followers",
		Method: http.MethodGet,
		Function: controllers.ListFollowers,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}/following",
		Method: http.MethodGet,
		Function: controllers.ListFollowing,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}/update-password",
		Method: http.MethodPost,
		Function: controllers.UpdatePassword,
		AuthRequired: true,
	},
}