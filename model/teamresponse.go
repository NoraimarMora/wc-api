package model

type TeamResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Confederation string    `json:"confederation"`
	FlagURl       string    `json:"flag_uri"`
	Group         string    `json:"group"`
	WorldRanking  int       `json:"world_ranking"`
	Appearances   int       `json:"appearances"`
	Host          bool      `json:"host"`
	Colors        Colors    `json:"colors"`
	Players       []*Player `json:"players"`
}
