package model

import "strings"

type Mascot struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	ImageURL    string   `json:"image_url"`
	Description []string `json:"description"`
}

type Mascots map[int64]*Mascot

func (m Mascots) ToSlice() []*Mascot {
	slice := make([]*Mascot, 0, len(m))
	for _, v := range m {
		slice = append(slice, v)
	}

	return slice
}

func (m Mascots) ByName(name string) Mascots {
	if name == "" {
		return m
	}

	mascots := make(Mascots)
	for k, v := range m {
		if strings.Contains(v.Name, name) {
			mascots[k] = v
		}
	}

	return mascots
}

func (m Mascots) ByCountry(country string) Mascots {
	if country == "" {
		return m
	}

	mascots := make(Mascots)
	for k, v := range m {
		if v.Country == country {
			mascots[k] = v
		}
	}

	return mascots
}
