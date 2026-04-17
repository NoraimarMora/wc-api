package chronology

import (
	"wc-api/model"
)

type Chronology struct {
	data model.ChronologyByMatch
}

func New(data model.ChronologyByMatch) *Chronology {
	return &Chronology{
		data: data,
	}
}

func (m *Chronology) GetByMatch(matchID int64) (*model.Chronology, error) {
	chronology, ok := m.data[matchID]
	if !ok {
		return nil, model.ErrNotFound
	}

	return chronology, nil
}
