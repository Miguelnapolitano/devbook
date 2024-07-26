package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Routes is a struct for all routes in API
type Routes struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Configure sets all routes inside of router
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function) .Methods(route.Method)
	}

	return r
}