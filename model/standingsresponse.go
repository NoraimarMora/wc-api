package model

type StandingsResponse struct {
	Team           *TeamResponse `json:"team"`
	Position       int           `json:"position"`
	Matches        int           `json:"matches"`
	Wins           int           `json:"wins"`
	Draws          int           `json:"draws"`
	Loss           int           `json:"loss"`
	GoalsScored    int           `json:"goals_scored"`
	GoalsAgainst   int           `json:"goals_against"`
	GoalDifference int           `json:"goal_difference"`
	Points         int           `json:"points"`
	Group          string        `json:"group"`
}

type StandingsByGroupResponse map[string][]*StandingsResponse
