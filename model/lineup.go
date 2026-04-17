package model

type LineUp struct {
	Formation       string    `json:"formation"`
	Coach           string    `json:"coach"`
	StartingPlayers []*Player `json:"starting_players"`
	Substitutes     []*Player `json:"substitutes"`
}

type LineUps map[string]*LineUp

// LineUpsByMatch ej.: 1: "home": {}, "away": {}
type LineUpsByMatch map[int64]LineUps
