package model

type City struct {
	ID          int64     `json:"id"`
	Country     string    `json:"country"`
	Name        string    `json:"name"`
	Description []string  `json:"description"`
	LogoURL     string    `json:"logo_url"`
	VideoURL    string    `json:"video_url"`
	ImageURL    string    `json:"image_url"`
	Stadium     Stadium   `json:"stadium"`
	ExtraInfo   ExtraInfo `json:"extra_info"`
	URL         string    `json:"url"`
}

type Stadium struct {
	Coordinates Coordinates `json:"coordinates"`
	Name        string      `json:"name"`
	ImageURL    string      `json:"image_url"`
	Capacity    int         `json:"capacity"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ExtraInfo struct {
	Title       string   `json:"title"`
	Description []string `json:"description"`
	Hashtag     string   `json:"hashtag"`
}

type Cities map[int64]*City

func (c Cities) ToSlice() []*City {
	slice := make([]*City, 0, len(c))
	for _, v := range c {
		slice = append(slice, v)
	}

	return slice
}

func (c Cities) ByCountry(country string) Cities {
	if country == "" {
		return c
	}

	cities := make(Cities)
	for k, v := range c {
		if v.Country == country {
			cities[k] = v
		}
	}

	return cities
}
