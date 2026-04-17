package local

import (
	"wc-api/infraestructure/repository/local/ball"
	"wc-api/infraestructure/repository/local/chronology"
	"wc-api/infraestructure/repository/local/city"
	"wc-api/infraestructure/repository/local/event"
	"wc-api/infraestructure/repository/local/lineup"
	"wc-api/infraestructure/repository/local/mascot"
	"wc-api/infraestructure/repository/local/match"
	"wc-api/infraestructure/repository/local/news"
	"wc-api/infraestructure/repository/local/player"
	"wc-api/infraestructure/repository/local/ranking"
	"wc-api/infraestructure/repository/local/record"
	"wc-api/infraestructure/repository/local/sound"
	"wc-api/infraestructure/repository/local/standings"
	"wc-api/infraestructure/repository/local/statistics"
	"wc-api/infraestructure/repository/local/team"
)

type Providers struct {
	Ball       *ball.Ball
	Chronology *chronology.Chronology
	City       *city.City
	Event      *event.Event
	LineUp     *lineup.LineUp
	Mascot     *mascot.Mascot
	Match      *match.Match
	News       *news.News
	Player     *player.Player
	Ranking    *ranking.Ranking
	Record     *record.Record
	Sound      *sound.Sound
	Standings  *standings.Standings
	Statistics *statistics.Statistics
	Team       *team.Team
}
