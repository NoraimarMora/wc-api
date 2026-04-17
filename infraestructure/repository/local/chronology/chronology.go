package chronology

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameChronologies = "chronologies.json"

type Chronology struct {
	common common.Common
}

func New() *Chronology {
	return &Chronology{
		common: common.New(),
	}
}

func (c *Chronology) GetData(ctx context.Context) (model.ChronologyByMatch, error) {
	file, err := c.common.GetFile(ctx, fileNameChronologies)
	if err != nil {
		return nil, err
	}

	var data model.ChronologyByMatch
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
