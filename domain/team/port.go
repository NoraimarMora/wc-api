package team

import "wc-api/model"

type Dependencies struct {
	Repository    Repository
	PlayerUseCase UseCasePlayer
}

type Repository interface {
	GetTeams(filters model.Filters) []*model.Team
	GetByID(id string) (*model.Team, error)
}

type UseCasePlayer interface {
	GetByTeam(teamID string) ([]*model.Player, error)
}
