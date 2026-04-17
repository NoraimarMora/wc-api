package lineup

import "wc-api/model"

type Repository interface {
	GetByMatch(matchID int64) (model.LineUps, error)
}

type LineUp struct {
	repository Repository
}

func New(repository Repository) *LineUp {
	return &LineUp{
		repository: repository,
	}
}

func (l *LineUp) GetByMatch(matchID int64) (model.LineUps, error) {
	return l.repository.GetByMatch(matchID)
}
