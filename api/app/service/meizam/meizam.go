package meizam

import (
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"
	"github.com/eyalkenig/meizam-api/api/app/service"
)

type Meizam struct {
	commandService service.Command
}

func NewMeizam(command service.Command) *Meizam {
	return &Meizam{commandService: command}
}

func (service *Meizam) CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error) {
	return service.commandService.CreateTeam(teamName, externalEntityId, imageUrl)
}
