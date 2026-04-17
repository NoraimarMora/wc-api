package match

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameMatches = "matches.json"

type Match struct {
	common common.Common
}

func New() *Match {
	return &Match{
		common: common.New(),
	}
}

func (m *Match) GetData(ctx context.Context) (model.Matches, error) {
	file, err := m.common.GetFile(ctx, fileNameMatches)
	if err != nil {
		return nil, err
	}

	var data model.Matches
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
