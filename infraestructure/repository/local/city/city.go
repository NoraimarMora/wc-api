package city

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameBall = "cities.json"

type City struct {
	common common.Common
}

func New() *City {
	return &City{
		common: common.New(),
	}
}

func (c *City) GetData(ctx context.Context) (model.Cities, error) {
	file, err := c.common.GetFile(ctx, fileNameBall)
	if err != nil {
		return nil, err
	}

	var data model.Cities
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
