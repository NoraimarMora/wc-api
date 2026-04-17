package ranking

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameRanking = "ranking.json"

type Ranking struct {
	common common.Common
}

func New() *Ranking {
	return &Ranking{
		common: common.New(),
	}
}

func (r *Ranking) GetData(ctx context.Context) (model.WorldRanking, error) {
	file, err := r.common.GetFile(ctx, fileNameRanking)
	if err != nil {
		return nil, err
	}

	var data model.WorldRanking
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
