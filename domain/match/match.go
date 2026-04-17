package match

import "wc-api/model"

type Match struct {
	repository Repository
	chronology UseCaseChronology
	lineUp     UseCaseLineUp
	statistics UseCaseStatistics
	city       UseCaseCity
	team       UseCaseTeam
}

func New(repository Repository, dependencies *Dependencies) *Match {
	return &Match{
		repository: repository,
		chronology: dependencies.ChronologyUseCase,
		lineUp:     dependencies.LineUpUseCase,
		statistics: dependencies.StatisticsUseCase,
		city:       dependencies.CityUseCase,
		team:       dependencies.TeamUseCase,
	}
}

func (m *Match) GetMatches(filters model.Filters) []*model.Match {
	return m.repository.GetMatches(filters)
}

func (m *Match) GetByID(id int64) (*model.MatchResponse, error) {
	match, err := m.repository.GetByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}

	chronology, err := m.chronology.GetByMatch(match.ID)
	if err != nil {
		chronology = &model.Chronology{}
	}

	lineUps, err := m.lineUp.GetByMatch(match.ID)
	if err != nil {
		lineUps = model.LineUps{
			"home": &model.LineUp{},
			"away": &model.LineUp{},
		}
	}

	statistics, err := m.statistics.GetByMatch(match.ID)
	if err != nil {
		statistics = make([]*model.Statistics, 0)
	}

	city, err := m.city.GetByID(match.CityID)
	if err != nil {
		city = &model.City{}
	}

	homeTeam, err := m.team.GetByID(match.HomeID)
	if err != nil {
		homeTeam = &model.TeamResponse{}
	}

	awayTeam, err := m.team.GetByID(match.AwayID)
	if err != nil {
		awayTeam = &model.TeamResponse{}
	}

	return &model.MatchResponse{
		ID:         match.ID,
		Date:       match.Date,
		Time:       match.Time,
		Status:     match.Status,
		Round:      match.Round,
		Group:      match.Group,
		Referee:    match.Referee,
		City:       city,
		HomeTeam:   homeTeam,
		AwayTeam:   awayTeam,
		HomeScore:  match.HomeScore,
		AwayScore:  match.AwayScore,
		LineUps:    lineUps,
		Statistics: statistics,
		Chronology: chronology,
		Highlights: match.Highlights,
	}, nil
}
