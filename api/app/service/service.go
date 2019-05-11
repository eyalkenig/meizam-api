package service

import "github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"

type Service interface {
	Command
	Query
}

type Command interface {
	CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error)
}

type Query interface {
	ListTeams(limit, offset int) ([]*models.Team, error)
}
