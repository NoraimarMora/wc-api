package sound

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNameSound = "sound.json"

type Sound struct {
	common common.Common
}

func New() *Sound {
	return &Sound{
		common: common.New(),
	}
}

func (s *Sound) GetData(ctx context.Context) (*model.Sound, error) {
	file, err := s.common.GetFile(ctx, fileNameSound)
	if err != nil {
		return nil, err
	}

	var data *model.Sound
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
