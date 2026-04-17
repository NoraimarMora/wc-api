package city

import "wc-api/model"

type Repository interface {
	GetCities(filters model.Filters) model.Cities
	GetByID(id int64) (*model.City, error)
}

type City struct {
	repository Repository
}

func New(repository Repository) *City {
	return &City{
		repository: repository,
	}
}

func (c *City) GetCities(filters model.Filters) model.Cities {
	return c.repository.GetCities(filters)
}

func (c *City) GetByID(id int64) (*model.City, error) {
	return c.repository.GetByID(id)
}
