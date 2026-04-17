package model

type Chronology struct {
	Substitutions []Substitution `json:"substitutions"`
	Fouls         []Foul         `json:"fouls"`
	Goals         []Goal         `json:"goals"`
}

type Substitution struct {
	PlayerIn  Player `json:"player_in"`
	PlayerOut Player `json:"player_out"`
	Time      int    `json:"time"`
}

type Foul struct {
	Player Player `json:"player"`
	Time   int    `json:"time"`
	Card   string `json:"type"`
}

type Goal struct {
	Player Player `json:"player"`
	Time   int    `json:"time"`
}

type ChronologyByMatch map[int64]*Chronology
