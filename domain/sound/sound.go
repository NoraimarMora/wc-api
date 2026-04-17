package sound

import (
	"wc-api/model"
)

type Repository interface {
	GetSound() *model.Sound
}

type Sound struct {
	repository Repository
}

func New(repository Repository) *Sound {
	return &Sound{
		repository: repository,
	}
}

func (s *Sound) GetSound() *model.Sound {
	return s.repository.GetSound()
}
