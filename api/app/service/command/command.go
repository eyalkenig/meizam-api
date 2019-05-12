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

func (command *MeizamCommand) CreateCompetition(name, competitionType string, externalEntityId *string) (*models.Competition, error) {
	if name == "" {
		return nil, fmt.Errorf("name is mandatory")
	}
	if competitionType == "" {
		return nil, fmt.Errorf("type is mandatory")
	}
	return command.repository.CreateCompetition(name, competitionType, externalEntityId)
}
