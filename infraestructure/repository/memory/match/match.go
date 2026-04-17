package match

import (
	"wc-api/model"
)

type Match struct {
	data model.Matches
}

func New(data model.Matches) *Match {
	return &Match{
		data: data,
	}
}

func (m *Match) GetMatches(filters model.Filters) []*model.Match {
	return m.data.
		ByCity(filters.CityID).
		ByGroup(filters.Group).
		ByRound(filters.Round).
		ByStatus(filters.Status).
		ByHomeID(filters.HomeID).
		ByAwayID(filters.AwayID).
		FromDate(filters.From).
		ToDate(filters.To).
		ToSlice()
}

func (m *Match) GetByID(id int64) (*model.Match, error) {
	match, ok := m.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}

	return match, nil
}
