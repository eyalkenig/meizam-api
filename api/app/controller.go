package app

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/eyalkenig/meizam-api/api/app/service"
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

type ErrorMessage struct {
	error string `json:"error"`
}

type Team struct {
	ID               int     `json:id`
	Name             string  `json:"name"`
	ExternalEntityId *string `json:"external_entity_id"`
	ImageUrl         *string `json:"image_url"`
}

func (controller *Controller) Ping(w http.ResponseWriter, req *http.Request) {
	test := &Ping{Pong: "ping pong"}
	encodeResponse(w, test)
}

func (controller *Controller) CreateTeam(w http.ResponseWriter, req *http.Request) {
	var teamRequest Team
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		handleServerError(w, err)
		return
	}
	if err := req.Body.Close(); err != nil {
		handleServerError(w, err)
		return
	}
	if err := json.Unmarshal(body, &teamRequest); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("Error CreateTeam unmarshalling data: %s", err.Error())
			handleServerError(w, err)
			return
		}
	}
	team, err := controller.service.CreateTeam(teamRequest.Name, teamRequest.ExternalEntityId, teamRequest.ImageUrl)

	if err != nil {
		handleServerError(w, err)
		return
	}

	encodeResponse(w, team)
}

func (controller *Controller) ListTeams(w http.ResponseWriter, req *http.Request) {
	pageUrlParam := getUrlParamWithDefault(req, "page", "0")
	page, err := strconv.Atoi(pageUrlParam)
	if err != nil || page < 0 {
		handleServerError(w, errors.New("page must be a positive number"))
		return
	}
	perPageParam := getUrlParamWithDefault(req, "per_page", "10")
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil || page < 0 {
		handleServerError(w, errors.New("per_page must be a positive number"))
		return
	}
	offset := page * perPage
	teams, err := controller.service.ListTeams(perPage, offset)

	if err != nil {
		handleServerError(w, err)
		return
	}

	encodeResponse(w, teams)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}

func handleServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	errorMessage := &ErrorMessage{error: err.Error()}
	encodedError := json.NewEncoder(w).Encode(errorMessage)
	if encodedError != nil {
		log.Printf("encoded with error: %s", err.Error())
		log.Printf("failed to encode with error: %s original error: %s", encodedError.Error(), err)
	}
}

func getUrlParamWithDefault(req *http.Request, param, defaultValue string) string {
	params := req.URL.Query()
	queryParam, ok := params[param]
	if ok && len(queryParam[0]) > 0 {
		return queryParam[0]
	}
	return defaultValue
}
