package standings

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameStandings = "standings.json"

type Standings struct {
	common common.Common
}

func New() *Standings {
	return &Standings{
		common: common.New(),
	}
}

func (s *Standings) GetData(ctx context.Context) (model.StandingsByGroup, error) {
	file, err := s.common.GetFile(ctx, fileNameStandings)
	if err != nil {
		return nil, err
	}

	var data model.StandingsByGroup
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
