package model

type UseCaseBall interface {
	GetBall() *Ball
}

type UseCaseChronology interface {
	GetByMatch(matchID int64) (*Chronology, error)
}

type UseCaseCity interface {
	GetCities(filters Filters) Cities
	GetByID(id int64) (*City, error)
}

type UseCaseEvent interface {
	GetAll() Events
}

type UseCaseLineUp interface {
	GetByMatch(matchID int64) (LineUps, error)
}

type UseCaseMascot interface {
	GetMascots(filters Filters) []*Mascot
	GetByID(id int64) (*Mascot, error)
}

type UseCaseMatch interface {
	GetMatches(filters Filters) []*Match
	GetByID(id int64) (*MatchResponse, error)
}

type UseCaseNews interface {
	GetAll() []*News
	GetByID(id int64) (*News, error)
}

type UseCasePlayer interface {
	GetByTeam(teamID string) ([]*Player, error)
}

type UseCaseRanking interface {
	GetWorldRanking() WorldRankingResponse
}

type UseCaseRecord interface {
	GetAll() []*Record
	GetByID(id int64) (*Record, error)
}

type UseCaseSound interface {
	GetSound() *Sound
}

type UseCaseStandings interface {
	GetAll() StandingsByGroupResponse
	GetByGroup(group string) ([]*StandingsResponse, error)
}

type UseCaseStatistics interface {
	GetByMatch(matchID int64) ([]*Statistics, error)
}

type UseCaseTeam interface {
	GetTeams(filters Filters) []*Team
	GetByID(id string) (*TeamResponse, error)
}

type Dependencies struct {
	BallUseCase       UseCaseBall
	ChronologyUseCase UseCaseChronology
	CityUseCase       UseCaseCity
	EventUseCase      UseCaseEvent
	LineUpUseCase     UseCaseLineUp
	MascotUseCase     UseCaseMascot
	MatchUseCase      UseCaseMatch
	NewsUseCase       UseCaseNews
	PlayerUseCase     UseCasePlayer
	RankingUseCase    UseCaseRanking
	RecordUseCase     UseCaseRecord
	SoundUseCase      UseCaseSound
	StandingsUseCase  UseCaseStandings
	StatisticsUseCase UseCaseStatistics
	TeamUseCase       UseCaseTeam
}
