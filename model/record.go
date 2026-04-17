package model

type Record struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	VideoSrc    string `json:"video_src"`
}

type Records map[int64]*Record

func (r Records) ToSlice() []*Record {
	slice := make([]*Record, 0, len(r))
	for _, v := range r {
		slice = append(slice, v)
	}

	return slice
}
