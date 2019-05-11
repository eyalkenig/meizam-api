package app

import (
	"log"

	"github.com/eyalkenig/meizam-api/api/app/controller"
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql"
	"github.com/eyalkenig/meizam-api/api/app/service/command"
	"github.com/eyalkenig/meizam-api/api/app/service/meizam"
	"github.com/eyalkenig/meizam-api/api/app/service/query"
)

type Container struct{}

func NewContainer() *Container {
	return &Container{}
}

func (container *Container) Resolve() (*Application, error) {
	mysqlBuilder := mysql.NewBuilder()
	repository, err := mysqlBuilder.Build()
	if err != nil {
		log.Fatal("Failed creating a new MySQL instance")
		return nil, err
	}

	commandService := command.NewMeizamCommand(repository)
	queryService := query.NewMeizamQuery(repository)
	service := meizam.NewMeizam(commandService, queryService)
	ctrl := controller.NewController(service)
	router := NewRouter(ctrl)

	return NewApplication(router), nil
}
