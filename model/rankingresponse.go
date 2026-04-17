package model

type RankingResponse struct {
	Team           *TeamResponse `json:"team"`
	Rank           int           `json:"rank"`
	PreviousRank   int           `json:"previous_rank"`
	Points         float64       `json:"points"`
	PreviousPoints float64       `json:"previous_points"`
}

type WorldRankingResponse []*RankingResponse
