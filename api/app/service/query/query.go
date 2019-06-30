package query

import (
	"github.com/eyalkenig/meizam-api/api/app/repository"
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"
)

func NewMeizamQuery(repository repository.Repository) *MeizamQuery {
	return &MeizamQuery{repository}
}

type MeizamQuery struct {
	repository repository.Repository
}

func (query *MeizamQuery) ListTeams(limit, offset int) ([]*models.Team, error) {
	return query.repository.ListTeams(limit, offset)
}

func (query *MeizamQuery) ListCompetitions(limit, offset int) ([]*models.Competition, error) {
	return query.repository.ListCompetitions(limit, offset)
}
