package repository

import "github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"

type Repository interface {
	realWorldRepository
}

type realWorldRepository interface {
	CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error)
	CreateCompetition(name, competitionType string, externalEntityId *string) (*models.Competition, error)
	ListTeams(limit, offset int) ([]*models.Team, error)
	ListCompetitions(limit, offset int) ([]*models.Competition, error)
}
