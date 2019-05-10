package command

import "github.com/eyalkenig/meizam-api/api/app/repository"

func NewMeizamCommand(repository repository.Repository) *MeizamCommand {
	return &MeizamCommand{repository}
}

type MeizamCommand struct {
	repository repository.Repository
}

func (command *MeizamCommand) CreateTeam(teamName, externalEntityId, imageUrl string) (int, error) {
	return -1, nil
}
