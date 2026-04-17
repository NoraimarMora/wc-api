package record

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameRecords = "records.json"

type Record struct {
	common common.Common
}

func New() *Record {
	return &Record{
		common: common.New(),
	}
}

func (r *Record) GetData(ctx context.Context) (model.Records, error) {
	file, err := r.common.GetFile(ctx, fileNameRecords)
	if err != nil {
		return nil, err
	}

	var data model.Records
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
