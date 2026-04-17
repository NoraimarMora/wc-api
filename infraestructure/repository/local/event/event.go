package event

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameEvents = "events.json"

type Event struct {
	common common.Common
}

func New() *Event {
	return &Event{
		common: common.New(),
	}
}

func (e *Event) GetData(ctx context.Context) (model.Events, error) {
	file, err := e.common.GetFile(ctx, fileNameEvents)
	if err != nil {
		return nil, err
	}

	var data model.Events
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
