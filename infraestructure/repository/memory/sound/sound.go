package sound

import (
	"wc-api/model"
)

type Sound struct {
	data *model.Sound
}

func New(data *model.Sound) *Sound {
	return &Sound{
		data: data,
	}
}

func (s *Sound) GetSound() *model.Sound {
	return s.data
}
