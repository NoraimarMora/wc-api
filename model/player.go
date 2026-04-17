package model

type Player struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Number   int    `json:"number"`
	PhotoURL int    `json:"photo_url"`
}

type PlayersByTeam map[string][]*Player
