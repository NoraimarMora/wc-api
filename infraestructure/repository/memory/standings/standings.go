package standings

import (
	"wc-api/model"
)

type Standings struct {
	data model.StandingsByGroup
}

func New(data model.StandingsByGroup) *Standings {
	return &Standings{
		data: data,
	}
}

func (s *Standings) GetAll() model.StandingsByGroup {
	return s.data
}

func (s *Standings) GetByGroup(group string) ([]*model.Standings, error) {
	standings, ok := s.data[group]
	if !ok {
		return nil, model.ErrNotFound
	}

	return standings, nil
}
