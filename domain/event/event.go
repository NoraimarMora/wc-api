package event

import "wc-api/model"

type Repository interface {
	GetAll() model.Events
}

type Event struct {
	repository Repository
}

func New(repository Repository) *Event {
	return &Event{
		repository: repository,
	}
}

func (e *Event) GetAll() model.Events {
	return e.repository.GetAll()
}
