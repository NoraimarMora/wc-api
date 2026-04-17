package model

type Standings struct {
	TeamID         string `json:"team_id"`
	Position       int    `json:"position"`
	Matches        int    `json:"matches"`
	Wins           int    `json:"wins"`
	Draws          int    `json:"draws"`
	Loss           int    `json:"loss"`
	GoalsScored    int    `json:"goals_scored"`
	GoalsAgainst   int    `json:"goals_against"`
	GoalDifference int    `json:"goal_difference"`
	Points         int    `json:"points"`
	Group          string `json:"group"`
}

type StandingsByGroup map[string][]*Standings
