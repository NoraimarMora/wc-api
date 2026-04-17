package player

import (
	"context"
	"encoding/json"
	"fmt"

	"wc-api/infraestructure/repository/local/common"
	"wc-api/model"
)

const fileNamePlayers = "players.json"

type Player struct {
	common common.Common
}

func New() *Player {
	return &Player{
		common: common.New(),
	}
}

func (p *Player) GetData(ctx context.Context) (model.PlayersByTeam, error) {
	file, err := p.common.GetFile(ctx, fileNamePlayers)
	if err != nil {
		return nil, err
	}

	var data model.PlayersByTeam
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %w", err)
	}

	return data, nil
}
