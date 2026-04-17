package statistics

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameStatistics = "statistics.json"

type Statistics struct {
	common common.Common
}

func New() *Statistics {
	return &Statistics{
		common: common.New(),
	}
}

func (s *Statistics) GetData(ctx context.Context) (model.StatisticsByMatch, error) {
	file, err := s.common.GetFile(ctx, fileNameStatistics)
	if err != nil {
		return nil, err
	}

	var data model.StatisticsByMatch
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
