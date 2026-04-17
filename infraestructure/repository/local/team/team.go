package team

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameTeams = "teams.json"

type Team struct {
	common common.Common
}

func New() *Team {
	return &Team{
		common: common.New(),
	}
}

func (t *Team) GetData(ctx context.Context) (model.Teams, error) {
	file, err := t.common.GetFile(ctx, fileNameTeams)
	if err != nil {
		return nil, err
	}

	var data model.Teams
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
