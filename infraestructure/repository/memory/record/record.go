package record

import (
	"wc-api/model"
)

type Record struct {
	data model.Records
}

func New(data model.Records) *Record {
	return &Record{
		data: data,
	}
}

func (r *Record) GetAll() []*model.Record {
	return r.data.ToSlice()
}

func (r *Record) GetByID(id int64) (*model.Record, error) {
	record, ok := r.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}

	return record, nil
}
