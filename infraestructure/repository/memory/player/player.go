package player

import "wc-api/model"

type Player struct {
	data model.PlayersByTeam
}

func New(data model.PlayersByTeam) *Player {
	return &Player{
		data: data,
	}
}

func (s *Player) GetByTeam(teamID string) ([]*model.Player, error) {
	players, ok := s.data[teamID]
	if !ok {
		return nil, model.ErrNotFound
	}

	return players, nil
}
