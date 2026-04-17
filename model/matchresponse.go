package model

type MatchResponse struct {
	ID         int64         `json:"id"`
	Date       string        `json:"date"`
	Time       string        `json:"time"`
	Status     string        `json:"status"`
	Round      int           `json:"round"`
	Group      string        `json:"group"`
	Referee    string        `json:"referee"`
	City       *City         `json:"city"`
	HomeTeam   *TeamResponse `json:"home_team"`
	AwayTeam   *TeamResponse `json:"away_team"`
	HomeScore  Score         `json:"home_score"`
	AwayScore  Score         `json:"away_score"`
	LineUps    LineUps       `json:"line_ups"`
	Statistics []*Statistics `json:"statistics"`
	Chronology *Chronology   `json:"chronology"`
	Highlights []Highlight   `json:"highlight"`
}
