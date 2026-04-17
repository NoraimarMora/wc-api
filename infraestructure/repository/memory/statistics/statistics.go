package statistics

import (
	"wc-api/model"
)

type Statistics struct {
	data model.StatisticsByMatch
}

func New(data model.StatisticsByMatch) *Statistics {
	return &Statistics{
		data: data,
	}
}

func (s *Statistics) GetByMatch(matchID int64) ([]*model.Statistics, error) {
	statistics, ok := s.data[matchID]
	if !ok {
		return nil, model.ErrNotFound
	}

	return statistics, nil
}
