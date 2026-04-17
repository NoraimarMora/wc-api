package ranking

import "wc-api/model"

type Ranking struct {
	repository Repository
	team       UseCaseTeam
}

func New(repository Repository, dependencies *Dependencies) *Ranking {
	return &Ranking{
		repository: repository,
		team:       dependencies.TeamUseCase,
	}
}

func (r *Ranking) GetWorldRanking() model.WorldRankingResponse {
	worldRanking := r.repository.GetWorldRanking()

	result := make(model.WorldRankingResponse, 0)
	for _, ranking := range worldRanking {
		team, err := r.team.GetByID(ranking.TeamID)
		if err != nil {
			continue
		}

		result = append(result, &model.RankingResponse{
			Team:           team,
			Rank:           ranking.Rank,
			PreviousRank:   ranking.PreviousRank,
			Points:         ranking.Points,
			PreviousPoints: ranking.PreviousPoints,
		})
	}

	return result
}
