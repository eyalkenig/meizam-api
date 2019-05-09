package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct{}

type Test struct {
	Name string `json:"name"`
}

func (controller *Controller) Test(w http.ResponseWriter, req *http.Request) {
	log.Println("im in controller")
	test := &Test{Name: "a test"}
	json.NewEncoder(w).Encode(test)
}
