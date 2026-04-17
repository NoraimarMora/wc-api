package player

import "wc-api/model"

type Repository interface {
	GetByTeam(teamID string) ([]*model.Player, error)
}

type Player struct {
	repository Repository
}

func New(repository Repository) *Player {
	return &Player{
		repository: repository,
	}
}

func (p *Player) GetByTeam(teamID string) ([]*model.Player, error) {
	return p.repository.GetByTeam(teamID)
}
