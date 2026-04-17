package record

import "wc-api/model"

type Repository interface {
	GetAll() []*model.Record
	GetByID(id int64) (*model.Record, error)
}

type Record struct {
	repository Repository
}

func New(repository Repository) *Record {
	return &Record{
		repository: repository,
	}
}

func (r *Record) GetAll() []*model.Record {
	return r.repository.GetAll()
}

func (r *Record) GetByID(id int64) (*model.Record, error) {
	return r.repository.GetByID(id)
}
