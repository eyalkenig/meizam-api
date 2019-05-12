package meizam

import (
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"
	"github.com/eyalkenig/meizam-api/api/app/service"
)

type Meizam struct {
	commandService service.Command
	queryService   service.Query
}

func NewMeizam(command service.Command, query service.Query) *Meizam {
	return &Meizam{commandService: command, queryService: query}
}

func (service *Meizam) CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error) {
	return service.commandService.CreateTeam(teamName, externalEntityId, imageUrl)
}

func (service *Meizam) ListTeams(limit, offset int) ([]*models.Team, error) {
	return service.queryService.ListTeams(limit, offset)
}
func (service *Meizam) CreateCompetition(name, competitionType string, externalEntityId *string) (*models.Competition, error) {
	return service.commandService.CreateCompetition(name, competitionType, externalEntityId)
}
