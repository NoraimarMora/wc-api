package model

type Statistics struct {
	Name      string  `json:"name"`
	Home      string  `json:"home"`
	Away      string  `json:"away"`
	HomeValue float64 `json:"home_value"`
	AwayValue float64 `json:"away_value"`
}

type StatisticsByMatch map[int64][]*Statistics
