package app

import (
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql"
	"github.com/eyalkenig/meizam-api/api/app/service/command"
	"github.com/eyalkenig/meizam-api/api/app/service/meizam"
	"log"
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
	service := meizam.NewMeizam(commandService)
	ctrl := NewController(service)
	router := NewRouter(ctrl)

	return NewApplication(router), nil
}
