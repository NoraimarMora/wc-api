package ball

import (
	"wc-api/model"
)

type Repository interface {
	GetBall() *model.Ball
}

type Ball struct {
	repository Repository
}

func New(repository Repository) *Ball {
	return &Ball{
		repository: repository,
	}
}

func (b *Ball) GetBall() *model.Ball {
	return b.repository.GetBall()
}
