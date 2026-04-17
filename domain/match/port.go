package match

import "wc-api/model"

type Dependencies struct {
	Repository        Repository
	ChronologyUseCase UseCaseChronology
	LineUpUseCase     UseCaseLineUp
	StatisticsUseCase UseCaseStatistics
	CityUseCase       UseCaseCity
	TeamUseCase       UseCaseTeam
}

type Repository interface {
	GetMatches(filters model.Filters) []*model.Match
	GetByID(id int64) (*model.Match, error)
}

type UseCaseChronology interface {
	GetByMatch(matchID int64) (*model.Chronology, error)
}

type UseCaseLineUp interface {
	GetByMatch(matchID int64) (model.LineUps, error)
}

type UseCaseStatistics interface {
	GetByMatch(matchID int64) ([]*model.Statistics, error)
}

type UseCaseCity interface {
	GetByID(id int64) (*model.City, error)
}

type UseCaseTeam interface {
	GetByID(id string) (*model.TeamResponse, error)
}
