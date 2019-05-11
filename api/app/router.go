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
			"ping",
			"GET",
			"/ping",
			controller.Ping,
		},
		Route{
			"create team",
			"POST",
			"/teams",
			controller.CreateTeam,
		},
		Route{
			"list team",
			"GET",
			"/teams",
			controller.ListTeams,
		},
	}
	return routes
}
