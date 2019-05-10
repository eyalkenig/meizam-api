package app

import "github.com/gorilla/mux"

type Application struct {
	Router *mux.Router
}

func NewApplication(router *mux.Router) *Application {
	return &Application{Router: router}
}
