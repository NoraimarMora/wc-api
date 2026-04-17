package news

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameNews = "news.json"

type News struct {
	common common.Common
}

func New() *News {
	return &News{
		common: common.New(),
	}
}

func (n *News) GetData(ctx context.Context) (model.NewsList, error) {
	file, err := n.common.GetFile(ctx, fileNameNews)
	if err != nil {
		return nil, err
	}

	var data model.NewsList
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
