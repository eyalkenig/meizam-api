package repository

import "github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"

type Repository interface {
	realWorldRepository
}

type realWorldRepository interface {
	CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error)
}
