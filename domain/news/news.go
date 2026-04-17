package news

import (
	"wc-api/model"
)

type Repository interface {
	GetAll() []*model.News
	GetByID(id int64) (*model.News, error)
}

type News struct {
	repository Repository
}

func New(repository Repository) *News {
	return &News{
		repository: repository,
	}
}

func (n *News) GetAll() []*model.News {
	return n.repository.GetAll()
}

func (n *News) GetByID(id int64) (*model.News, error) {
	return n.repository.GetByID(id)
}
