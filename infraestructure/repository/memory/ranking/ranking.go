package ranking

import (
	"wc-api/model"
)

type Ranking struct {
	data model.WorldRanking
}

func New(data model.WorldRanking) *Ranking {
	return &Ranking{
		data: data,
	}
}

func (r *Ranking) GetWorldRanking() model.WorldRanking {
	return r.data
}
