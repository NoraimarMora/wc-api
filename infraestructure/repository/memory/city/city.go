package city

import (
	"wc-api/model"
)

type City struct {
	data model.Cities
}

func New(data model.Cities) *City {
	return &City{
		data: data,
	}
}

func (c *City) GetCities(filters model.Filters) model.Cities {
	return c.data.ByCountry(filters.Country)
}

func (c *City) GetByID(id int64) (*model.City, error) {
	city, ok := c.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}

	return city, nil
}
