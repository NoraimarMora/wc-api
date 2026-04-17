package model

type Ranking struct {
	TeamID         string  `json:"team_id"`
	Rank           int     `json:"rank"`
	PreviousRank   int     `json:"previous_rank"`
	Points         float64 `json:"points"`
	PreviousPoints float64 `json:"previous_points"`
}

type WorldRanking []*Ranking
