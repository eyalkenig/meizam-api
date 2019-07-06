package app

import (
	"github.com/codegangsta/negroni"
	"log"
	"net/http"

	"github.com/eyalkenig/meizam-api/api/app/controller"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/mux"
)

type Route struct {
	Name          string
	Method        string
	Pattern       string
	HandlerFunc   http.HandlerFunc
	Authenticated bool
}

type Routes []Route

func NewRouter(controller *controller.Controller, middleware *jwtmiddleware.JWTMiddleware) *mux.Router {
	authRoutes := defineAuthenticatedRoutes(controller)
	router := mux.NewRouter().StrictSlash(true)
	for _, appRoute := range authRoutes {
		log.Println(appRoute.Name)

		route := router.
			Methods(appRoute.Method).
			Path(appRoute.Pattern).
			Name(appRoute.Name)
		if appRoute.Authenticated && middleware != nil {
			route.Handler(negroni.New(negroni.HandlerFunc(middleware.HandlerWithNext), negroni.Wrap(appRoute.HandlerFunc)))
		} else {
			route.Handler(appRoute.HandlerFunc)
		}
	}
	return router
}
func defineAuthenticatedRoutes(controller *controller.Controller) []Route {
	var routes = Routes{
		Route{
			"ping",
			"GET",
			"/ping",
			controller.Ping,
			false,
		},
		Route{
			"create team",
			"POST",
			"/v1/teams",
			controller.CreateTeam,
			true,
		},
		Route{
			"list team",
			"GET",
			"/v1/teams",
			controller.ListTeams,
			true,
		},
		Route{
			"create competition",
			"POST",
			"/v1/competitions",
			controller.CreateCompetition,
			true,
		},
		Route{
			"list competition",
			"GET",
			"/v1/competitions",
			controller.ListCompetitions,
			true,
		},
	}
	return routes
}
