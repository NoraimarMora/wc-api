package lineup

import (
	"wc-api/model"
)

type LineUp struct {
	data model.LineUpsByMatch
}

func New(data model.LineUpsByMatch) *LineUp {
	return &LineUp{
		data: data,
	}
}

func (l *LineUp) GetByMatch(matchID int64) (model.LineUps, error) {
	lineUps, ok := l.data[matchID]
	if !ok {
		return nil, model.ErrNotFound
	}

	return lineUps, nil
}
