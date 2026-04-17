package news

import (
	"wc-api/model"
)

type News struct {
	data model.NewsList
}

func New(data model.NewsList) *News {
	return &News{
		data: data,
	}
}

func (n *News) GetAll() []*model.News {
	return n.data.ToSlice()
}

func (n *News) GetByID(id int64) (*model.News, error) {
	news, ok := n.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}

	return news, nil
}
