package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, routesPublications...)

	for _, route := range routes {

		if route.AuthRequired {
			
			r.HandleFunc(route.URI, 
				middlewares.Logger(middlewares.Authenticates(route.Function)),
				).Methods(route.Method)
		}else{
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}