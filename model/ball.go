package model

type Ball struct {
	Name      string    `json:"name"`
	Features  []Feature `json:"features"`
	ImagesURL []string  `json:"images_url"`
}

type Feature struct {
	Title       string   `json:"title"`
	Description []string `json:"description"`
}
