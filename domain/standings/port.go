package standings

import "wc-api/model"

type Dependencies struct {
	Repository  Repository
	TeamUseCase UseCaseTeam
}

type Repository interface {
	GetAll() model.StandingsByGroup
	GetByGroup(group string) ([]*model.Standings, error)
}

type UseCaseTeam interface {
	GetByID(id string) (*model.TeamResponse, error)
}
