package app

import (
	"encoding/json"
	"github.com/eyalkenig/meizam-api/api/app/service"
	"net/http"
)

func NewController(service service.Service) *Controller {
	return &Controller{service}
}

type Controller struct {
	service service.Service
}

type Ping struct {
	Pong string `json:"pong"`
}

func (controller *Controller) Ping(w http.ResponseWriter, req *http.Request) {
	test := &Ping{Pong: "ping pong"}
	json.NewEncoder(w).Encode(test)
}
