package controller

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

type Team struct {
	ID               int     `json:id`
	Name             string  `json:"name"`
	ExternalEntityId *string `json:"external_entity_id"`
	ImageUrl         *string `json:"image_url"`
}

type Competition struct {
	ID               int     `json:id`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	ExternalEntityId *string `json:"external_entity_id"`
}

func (controller *Controller) Ping(w http.ResponseWriter, req *http.Request) {
	test := &Ping{Pong: "ping pong"}
	encodeResponse(w, test)
}

func (controller *Controller) CreateTeam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var teamRequest Team
	if err := getRequestBody(w, req, &teamRequest); err != nil {
		return
	}

	team, err := controller.service.CreateTeam(teamRequest.Name, teamRequest.ExternalEntityId, teamRequest.ImageUrl)

	if err != nil {
		handleServerError(w, http.StatusInternalServerError, err)
		return
	}
	encodeResponse(w, team)
}

func (controller *Controller) ListTeams(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pageUrlParam := getUrlParamWithDefault(req, "page", "0")
	page, err := strconv.Atoi(pageUrlParam)
	if err != nil || page < 0 {
		handleServerError(w, http.StatusBadRequest, errors.New("page must be a positive number"))
		return
	}
	perPageParam := getUrlParamWithDefault(req, "per_page", "10")
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil || perPage < 0 {
		handleServerError(w, http.StatusBadRequest, errors.New("per_page must be a positive number"))
		return
	}
	offset := page * perPage
	teams, err := controller.service.ListTeams(perPage, offset)

	if err != nil {
		handleServerError(w, http.StatusInternalServerError, err)
		return
	}

	encodeResponse(w, teams)
}

func (controller *Controller) CreateCompetition(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var competitionRequest Competition
	if err := getRequestBody(w, req, &competitionRequest); err != nil {
		return
	}

	competition, err := controller.service.CreateCompetition(competitionRequest.Name, competitionRequest.Type, competitionRequest.ExternalEntityId)

	if err != nil {
		handleServerError(w, http.StatusInternalServerError, err)
		return
	}
	encodeResponse(w, competition)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode with error: %s, response: %v", err.Error(), response)
	}
}

func handleServerError(w http.ResponseWriter, statusCode int, err error) {
	log.Printf("response error: %s", err.Error())
	w.WriteHeader(statusCode)
	result := map[string]string{}
	result["error"] = err.Error()
	encodeResponse(w, result)
}

func getUrlParamWithDefault(req *http.Request, param, defaultValue string) string {
	params := req.URL.Query()
	queryParam, ok := params[param]
	if ok && len(queryParam[0]) > 0 {
		return queryParam[0]
	}
	return defaultValue
}

func getRequestBody(w http.ResponseWriter, req *http.Request, value interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		handleServerError(w, http.StatusInternalServerError, err)
		return err
	}
	if err := req.Body.Close(); err != nil {
		handleServerError(w, http.StatusInternalServerError, err)
		return err
	}
	if err := json.Unmarshal(body, value); err != nil {
		handleServerError(w, http.StatusUnprocessableEntity, err)
		return err
	}
	return nil
}
