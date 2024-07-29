package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Routes{
	{
		URI:          "/publications",
		Method:       http.MethodPost,
		Function:     controllers.CreatePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications",
		Method:       http.MethodGet,
		Function:     controllers.ListPublications,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodGet,
		Function:     controllers.RetrievePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePublication,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/publications",
		Method:       http.MethodGet,
		Function:     controllers.ListUserPublications,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationId}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationId}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePublication,
		AuthRequired: true,
	},
}
