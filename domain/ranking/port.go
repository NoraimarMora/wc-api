package ranking

import "wc-api/model"

type Dependencies struct {
	Repository  Repository
	TeamUseCase UseCaseTeam
}

type Repository interface {
	GetWorldRanking() model.WorldRanking
}

type UseCaseTeam interface {
	GetByID(id string) (*model.TeamResponse, error)
}
