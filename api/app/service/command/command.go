package command

import (
	"fmt"

	"github.com/eyalkenig/meizam-api/api/app/repository"
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"
)

func NewMeizamCommand(repository repository.Repository) *MeizamCommand {
	return &MeizamCommand{repository}
}

type MeizamCommand struct {
	repository repository.Repository
}

func (command *MeizamCommand) CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error) {
	if teamName == "" {
		return nil, fmt.Errorf("team name is mandatory")
	}
	return command.repository.CreateTeam(teamName, externalEntityId, imageUrl)
}
