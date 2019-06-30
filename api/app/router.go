package app

import (
	"log"
	"net/http"

	"github.com/eyalkenig/meizam-api/api/app/controller"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(controller *controller.Controller) *mux.Router {
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

func defineRoutes(controller *controller.Controller) []Route {
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
		Route{
			"create competition",
			"POST",
			"/competitions",
			controller.CreateCompetition,
		},
		Route{
			"list competition",
			"GET",
			"/competitions",
			controller.ListCompetitions,
		},
	}
	return routes
}
