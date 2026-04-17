package team

import (
	"wc-api/model"
)

type Team struct {
	data model.Teams
}

func New(data model.Teams) *Team {
	return &Team{
		data: data,
	}
}

func (t *Team) GetTeams(filters model.Filters) []*model.Team {
	return t.data.
		ByGroup(filters.Group).
		ByConfederation(filters.Confederation).
		ToSlice()
}

func (t *Team) GetByID(id string) (*model.Team, error) {
	team, ok := t.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}

	return team, nil
}
