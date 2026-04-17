package chronology

import (
	"wc-api/model"
)

type Repository interface {
	GetByMatch(matchID int64) (*model.Chronology, error)
}

type Chronology struct {
	repository Repository
}

func New(repository Repository) *Chronology {
	return &Chronology{
		repository: repository,
	}
}

func (m *Chronology) GetByMatch(matchID int64) (*model.Chronology, error) {
	return m.repository.GetByMatch(matchID)
}
