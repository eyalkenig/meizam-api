package app

import (
	"encoding/json"
	"github.com/eyalkenig/meizam-api/api/app/service"
	"log"
	"net/http"
)

func NewController(service service.Service) *Controller {
	return &Controller{service}
}

type Controller struct {
	service service.Service
}

type Test struct {
	Name string `json:"name"`
}

func (controller *Controller) Test(w http.ResponseWriter, req *http.Request) {
	log.Println("im in controller")
	test := &Test{Name: "a test"}
	json.NewEncoder(w).Encode(test)
}
