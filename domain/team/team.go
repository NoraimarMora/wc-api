package team

import (
	"wc-api/model"
)

type Team struct {
	repository Repository
	player     UseCasePlayer
}

func New(repository Repository, dependencies *Dependencies) *Team {
	return &Team{
		repository: repository,
		player:     dependencies.PlayerUseCase,
	}
}

func (t *Team) GetTeams(filters model.Filters) []*model.Team {
	return t.repository.GetTeams(filters)
}

func (t *Team) GetByID(id string) (*model.TeamResponse, error) {
	team, err := t.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	players, err := t.player.GetByTeam(team.ID)
	if err != nil {
		return nil, err
	}

	return &model.TeamResponse{
		ID:            team.ID,
		Name:          team.Name,
		Confederation: team.Confederation,
		FlagURl:       team.FlagURl,
		Group:         team.Group,
		WorldRanking:  team.WorldRanking,
		Appearances:   team.Appearances,
		Host:          team.Host,
		Colors:        team.Colors,
		Players:       players,
	}, nil
}
