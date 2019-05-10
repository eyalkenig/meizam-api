package meizam

import (
	"github.com/eyalkenig/meizam-api/api/app/service"
)

type Meizam struct {
	commandService service.Command
}

func NewMeizam(command service.Command) *Meizam {
	return &Meizam{commandService: command}
}

func (service *Meizam) CreateTeam(teamName, externalEntityId, imageUrl string) (int, error) {
	return -1, nil
}
