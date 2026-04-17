package model

import "time"

type Filters struct {
	CityID        int64     `json:"city"`
	Round         int       `json:"round"`
	Status        string    `json:"status"`
	Group         string    `json:"group"`
	HomeID        string    `json:"home_id"`
	AwayID        string    `json:"away_id"`
	Country       string    `json:"country"`
	Name          string    `json:"name"`
	Confederation string    `json:"confederation"`
	From          time.Time `json:"from"`
	To            time.Time `json:"to"`
}
