package model

type Sound struct {
	Title    string    `json:"title"`
	Resume   string    `json:"resume"`
	Features []Feature `json:"features"`
	URL      string    `json:"url"`
	ImageURL string    `json:"image_url"`
}
