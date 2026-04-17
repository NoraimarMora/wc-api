package statistics

import "wc-api/model"

type Repository interface {
	GetByMatch(matchID int64) ([]*model.Statistics, error)
}

type Statistics struct {
	repository Repository
}

func New(repository Repository) *Statistics {
	return &Statistics{
		repository: repository,
	}
}

func (s *Statistics) GetByMatch(matchID int64) ([]*model.Statistics, error) {
	return s.repository.GetByMatch(matchID)
}
