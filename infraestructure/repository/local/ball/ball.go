package ball

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameBall = "ball.json"

type Ball struct {
	common common.Common
}

func New() *Ball {
	return &Ball{
		common: common.New(),
	}
}

func (b *Ball) GetData(ctx context.Context) (*model.Ball, error) {
	file, err := b.common.GetFile(ctx, fileNameBall)
	if err != nil {
		return nil, err
	}

	var data *model.Ball
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
