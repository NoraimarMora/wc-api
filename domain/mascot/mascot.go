package mascot

import (
	"wc-api/model"
)

type Repository interface {
	GetMascots(filters model.Filters) []*model.Mascot
	GetByID(id int64) (*model.Mascot, error)
}

type Mascot struct {
	repository Repository
}

func New(repository Repository) *Mascot {
	return &Mascot{
		repository: repository,
	}
}

func (m *Mascot) GetMascots(filters model.Filters) []*model.Mascot {
	return m.repository.GetMascots(filters)
}

func (m *Mascot) GetByID(id int64) (*model.Mascot, error) {
	return m.repository.GetByID(id)
}
