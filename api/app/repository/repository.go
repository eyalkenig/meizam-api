package repository

type Repository interface {
	realWorldRepository
}

type realWorldRepository interface {
	CreateTeam(teamName, externalEntityId, imageUrl string) (int, error)
}
