package standings

import "wc-api/model"

type Standings struct {
	repository  Repository
	teamUseCase UseCaseTeam
}

func New(repository Repository, dependencies *Dependencies) *Standings {
	return &Standings{
		repository:  repository,
		teamUseCase: dependencies.TeamUseCase,
	}
}

func (s *Standings) GetAll() model.StandingsByGroupResponse {
	standingsByGroup := s.repository.GetAll()

	result := make(model.StandingsByGroupResponse)
	for group, byGroup := range standingsByGroup {
		for _, standings := range byGroup {
			if _, ok := result[group]; !ok {
				result[group] = make([]*model.StandingsResponse, 0)
			}

			team, err := s.teamUseCase.GetByID(group)
			if err != nil {
				continue
			}

			result[group] = append(result[group], &model.StandingsResponse{
				Team:           team,
				Position:       standings.Position,
				Matches:        standings.Matches,
				Wins:           standings.Wins,
				Draws:          standings.Draws,
				Loss:           standings.Loss,
				GoalsScored:    standings.GoalsScored,
				GoalsAgainst:   standings.GoalsAgainst,
				GoalDifference: standings.GoalDifference,
				Points:         standings.Points,
				Group:          group,
			})
		}
	}

	return result
}

func (s *Standings) GetByGroup(group string) ([]*model.StandingsResponse, error) {
	standingsByGroup, err := s.repository.GetByGroup(group)
	if err != nil {
		return nil, err
	}

	result := make([]*model.StandingsResponse, 0)
	for _, standings := range standingsByGroup {
		team, err := s.teamUseCase.GetByID(group)
		if err != nil {
			continue
		}

		result = append(result, &model.StandingsResponse{
			Team:           team,
			Position:       standings.Position,
			Matches:        standings.Matches,
			Wins:           standings.Wins,
			Draws:          standings.Draws,
			Loss:           standings.Loss,
			GoalsScored:    standings.GoalsScored,
			GoalsAgainst:   standings.GoalsAgainst,
			GoalDifference: standings.GoalDifference,
			Points:         standings.Points,
			Group:          group,
		})
	}

	return result, nil
}
