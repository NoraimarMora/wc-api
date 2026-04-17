package local

import (
	"context"

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
	"wc-api/model"
)

type Local struct {
	providers Providers
}

func New() *Local {
	return &Local{
		providers: Providers{
			Ball:       ball.New(),
			Chronology: chronology.New(),
			City:       city.New(),
			Event:      event.New(),
			LineUp:     lineup.New(),
			Mascot:     mascot.New(),
			Match:      match.New(),
			News:       news.New(),
			Player:     player.New(),
			Ranking:    ranking.New(),
			Record:     record.New(),
			Sound:      sound.New(),
			Standings:  standings.New(),
			Statistics: statistics.New(),
			Team:       team.New(),
		},
	}
}

func (l *Local) GetProviders() Providers {
	return l.providers
}

func (l *Local) LoadBall(ctx context.Context) (*model.Ball, error) {
	return l.providers.Ball.GetData(ctx)
}

func (l *Local) LoadChronologyByMatch(ctx context.Context) (model.ChronologyByMatch, error) {
	return l.providers.Chronology.GetData(ctx)
}

func (l *Local) LoadCities(ctx context.Context) (model.Cities, error) {
	return l.providers.City.GetData(ctx)
}

func (l *Local) LoadEvents(ctx context.Context) (model.Events, error) {
	return l.providers.Event.GetData(ctx)
}

func (l *Local) LoadLineUpsByMatch(ctx context.Context) (model.LineUpsByMatch, error) {
	return l.providers.LineUp.GetData(ctx)
}

func (l *Local) LoadMascots(ctx context.Context) (model.Mascots, error) {
	return l.providers.Mascot.GetData(ctx)
}

func (l *Local) LoadMatches(ctx context.Context) (model.Matches, error) {
	return l.providers.Match.GetData(ctx)
}

func (l *Local) LoadNewsList(ctx context.Context) (model.NewsList, error) {
	return l.providers.News.GetData(ctx)
}

func (l *Local) LoadPlayersByTeam(ctx context.Context) (model.PlayersByTeam, error) {
	return l.providers.Player.GetData(ctx)
}

func (l *Local) LoadWorldRanking(ctx context.Context) (model.WorldRanking, error) {
	return l.providers.Ranking.GetData(ctx)
}

func (l *Local) LoadRecords(ctx context.Context) (model.Records, error) {
	return l.providers.Record.GetData(ctx)
}

func (l *Local) LoadSound(ctx context.Context) (*model.Sound, error) {
	return l.providers.Sound.GetData(ctx)
}

func (l *Local) LoadStandingsByGroup(ctx context.Context) (model.StandingsByGroup, error) {
	return l.providers.Standings.GetData(ctx)
}

func (l *Local) LoadStatisticsByMatch(ctx context.Context) (model.StatisticsByMatch, error) {
	return l.providers.Statistics.GetData(ctx)
}

func (l *Local) LoadTeams(ctx context.Context) (model.Teams, error) {
	return l.providers.Team.GetData(ctx)
}
