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
		AuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Function: controllers.RetrieveUser,
		AuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthRequired: false,
	},
}