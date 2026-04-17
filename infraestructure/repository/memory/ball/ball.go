package ball

import (
	"wc-api/model"
)

type Ball struct {
	data *model.Ball
}

func New(data *model.Ball) *Ball {
	return &Ball{
		data: data,
	}
}

func (b *Ball) GetBall() *model.Ball {
	return b.data
}
