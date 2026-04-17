package event

import (
	"wc-api/model"
)

type Event struct {
	data model.Events
}

func New(data model.Events) *Event {
	return &Event{
		data: data,
	}
}

func (e *Event) GetAll() model.Events {
	return e.data
}
