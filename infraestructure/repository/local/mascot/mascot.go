package mascot

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameMascot = "mascot.json"

type Mascot struct {
	common common.Common
}

func New() *Mascot {
	return &Mascot{
		common: common.New(),
	}
}

func (m *Mascot) GetData(ctx context.Context) (model.Mascots, error) {
	file, err := m.common.GetFile(ctx, fileNameMascot)
	if err != nil {
		return nil, err
	}

	var data model.Mascots
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
