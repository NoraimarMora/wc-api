package mascot

import (
	"wc-api/model"
)

type Mascot struct {
	data model.Mascots
}

func New(data model.Mascots) *Mascot {
	return &Mascot{
		data: data,
	}
}

func (m *Mascot) GetMascots(filters model.Filters) []*model.Mascot {
	return m.data.ByCountry(filters.Country).ByName(filters.Name).ToSlice()
}

func (m *Mascot) GetByID(id int64) (*model.Mascot, error) {
	mascot, ok := m.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}

	return mascot, nil
}
