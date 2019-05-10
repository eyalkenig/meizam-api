package app

import (
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(controller *Controller) *mux.Router {
	routes := defineRoutes(controller)
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func defineRoutes(controller *Controller) []Route {
	var routes = Routes{
		Route{
			"test",
			"GET",
			"/test",
			controller.Test,
		},
	}
	return routes
}
