package service

type Service interface {
	Command
	Query
}

type Command interface {
	CreateTeam(teamName, externalEntityId, imageUrl string) (int, error)
}

type Query interface {
}
