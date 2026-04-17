package lineup

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameLineUps = "lineups.json"

type LineUp struct {
	common common.Common
}

func New() *LineUp {
	return &LineUp{
		common: common.New(),
	}
}

func (t *LineUp) GetData(ctx context.Context) (model.LineUpsByMatch, error) {
	file, err := t.common.GetFile(ctx, fileNameLineUps)
	if err != nil {
		return nil, err
	}

	var data model.LineUpsByMatch
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
