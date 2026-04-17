package bootstrap

import (
	"context"

	"github.com/rs/zerolog/log"

	"wc-api/domain/ball"
	"wc-api/domain/chronology"
	"wc-api/domain/city"
	"wc-api/domain/event"
	"wc-api/domain/lineup"
	"wc-api/domain/mascot"
	"wc-api/domain/match"
	"wc-api/domain/news"
	"wc-api/domain/player"
	"wc-api/domain/ranking"
	"wc-api/domain/record"
	"wc-api/domain/sound"
	"wc-api/domain/standings"
	"wc-api/domain/statistics"
	"wc-api/domain/team"
	"wc-api/infraestructure/repository/local"
	ballRepository "wc-api/infraestructure/repository/memory/ball"
	chronologyRepository "wc-api/infraestructure/repository/memory/chronology"
	cityRepository "wc-api/infraestructure/repository/memory/city"
	eventRepository "wc-api/infraestructure/repository/memory/event"
	lineupRepository "wc-api/infraestructure/repository/memory/lineup"
	mascotRepository "wc-api/infraestructure/repository/memory/mascot"
	matchRepository "wc-api/infraestructure/repository/memory/match"
	newsRepository "wc-api/infraestructure/repository/memory/news"
	playerRepository "wc-api/infraestructure/repository/memory/player"
	rankingRepository "wc-api/infraestructure/repository/memory/ranking"
	recordRepository "wc-api/infraestructure/repository/memory/record"
	soundRepository "wc-api/infraestructure/repository/memory/sound"
	standingsRepository "wc-api/infraestructure/repository/memory/standings"
	statisticsRepository "wc-api/infraestructure/repository/memory/statistics"
	teamRepository "wc-api/infraestructure/repository/memory/team"
	"wc-api/model"
)

type Dependencies struct {
	appName string
}

func NewDependencies(ctx context.Context, appName string) Dependencies {
	return Dependencies{
		appName: appName,
	}
}

func (d *Dependencies) Build(ctx context.Context, provider *local.Local) model.Dependencies {
	ballInfo, err := provider.LoadBall(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading ball")
	}

	chronologies, err := provider.LoadChronologyByMatch(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading chronology")
	}

	cities, err := provider.LoadCities(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading cities")
	}

	events, err := provider.LoadEvents(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading events")
	}

	lineUpsByMatch, err := provider.LoadLineUpsByMatch(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading lineups")
	}

	mascots, err := provider.LoadMascots(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading mascots")
	}

	matches, err := provider.LoadMatches(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading matches")
	}

	newsList, err := provider.LoadNewsList(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading new list")
	}

	playersByTeam, err := provider.LoadPlayersByTeam(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading players")
	}

	worldRanking, err := provider.LoadWorldRanking(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading world ranking")
	}

	records, err := provider.LoadRecords(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading records")
	}

	soundInfo, err := provider.LoadSound(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading sound")
	}

	standingsByGroup, err := provider.LoadStandingsByGroup(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading standings")
	}

	statisticsByMatch, err := provider.LoadStatisticsByMatch(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading statistics")
	}

	teams, err := provider.LoadTeams(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("error loading teams")
	}

	ballUseCase := ball.New(ballRepository.New(ballInfo))
	chronologyUseCase := chronology.New(chronologyRepository.New(chronologies))
	cityUseCase := city.New(cityRepository.New(cities))
	eventUseCase := event.New(eventRepository.New(events))
	lineUpUseCase := lineup.New(lineupRepository.New(lineUpsByMatch))
	mascotUseCase := mascot.New(mascotRepository.New(mascots))
	newsUseCase := news.New(newsRepository.New(newsList))
	playerUseCase := player.New(playerRepository.New(playersByTeam))
	recordUseCase := record.New(recordRepository.New(records))
	soundUseCase := sound.New(soundRepository.New(soundInfo))
	statisticsUseCase := statistics.New(statisticsRepository.New(statisticsByMatch))
	teamUseCase := team.New(teamRepository.New(teams), &team.Dependencies{
		PlayerUseCase: playerUseCase,
	})
	matchUseCase := match.New(matchRepository.New(matches), &match.Dependencies{
		ChronologyUseCase: chronologyUseCase,
		LineUpUseCase:     lineUpUseCase,
		StatisticsUseCase: statisticsUseCase,
		CityUseCase:       cityUseCase,
		TeamUseCase:       teamUseCase,
	})
	standingsUseCase := standings.New(standingsRepository.New(standingsByGroup), &standings.Dependencies{
		TeamUseCase: teamUseCase,
	})
	rankingUseCase := ranking.New(rankingRepository.New(worldRanking), &ranking.Dependencies{
		TeamUseCase: teamUseCase,
	})

	return model.Dependencies{
		BallUseCase:       ballUseCase,
		ChronologyUseCase: chronologyUseCase,
		CityUseCase:       cityUseCase,
		EventUseCase:      eventUseCase,
		LineUpUseCase:     lineUpUseCase,
		MascotUseCase:     mascotUseCase,
		MatchUseCase:      matchUseCase,
		NewsUseCase:       newsUseCase,
		RankingUseCase:    rankingUseCase,
		RecordUseCase:     recordUseCase,
		SoundUseCase:      soundUseCase,
		StandingsUseCase:  standingsUseCase,
		StatisticsUseCase: statisticsUseCase,
		TeamUseCase:       teamUseCase,
	}
}
