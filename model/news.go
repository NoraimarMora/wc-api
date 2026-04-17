package model

type News struct {
	Title         string `json:"title"`
	PreviewText   string `json:"preview_text"`
	URL           string `json:"url"`
	ImageURL      string `json:"image_url"`
	ImageTitle    string `json:"image_title"`
	PublishedDate string `json:"published_date"`
}

type NewsList map[int64]*News

func (n NewsList) ToSlice() []*News {
	slice := make([]*News, 0, len(n))
	for _, v := range n {
		slice = append(slice, v)
	}

	return slice
}
