package model

import "time"

type Matches map[int64]*Match

type Match struct {
	ID         int64       `json:"id"`
	Date       string      `json:"date"`
	Time       string      `json:"time"`
	Status     string      `json:"status"`
	Round      int         `json:"round"`
	Group      string      `json:"group"`
	Referee    string      `json:"referee"`
	CityID     int64       `json:"city"`
	HomeID     string      `json:"home_id"`
	AwayID     string      `json:"away_id"`
	HomeScore  Score       `json:"home_score"`
	AwayScore  Score       `json:"away_score"`
	Highlights []Highlight `json:"highlight"`
}

type Score struct {
	Total        int `json:"total"`
	Period1      int `json:"period1"`
	Period2      int `json:"period2"`
	ExtraTime    int `json:"extra_time"`
	PenaltyRound int `json:"penalty"`
}

type Highlight struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

func (m Matches) ToSlice() []*Match {
	slice := make([]*Match, 0, len(m))
	for _, v := range m {
		slice = append(slice, v)
	}

	return slice
}

func (m Matches) ByCity(cityID int64) Matches {
	if cityID == 0 {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		if v.CityID == cityID {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) ByRound(round int) Matches {
	if round == 0 {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		if v.Round == round {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) ByStatus(status string) Matches {
	if status == "" {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		if v.Status == status {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) ByGroup(group string) Matches {
	if group == "" {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		if v.Group == group {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) ByHomeID(homeID string) Matches {
	if homeID == "" {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		if v.HomeID == homeID {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) ByAwayID(awayID string) Matches {
	if awayID == "" {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		if v.AwayID == awayID {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) FromDate(from time.Time) Matches {
	if from.IsZero() {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		matchDate, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			continue
		}

		if !matchDate.Before(from) {
			matches[k] = v
		}
	}

	return matches
}

func (m Matches) ToDate(to time.Time) Matches {
	if to.IsZero() {
		return m
	}

	matches := make(Matches)
	for k, v := range m {
		matchDate, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			continue
		}

		if !matchDate.After(to) {
			matches[k] = v
		}
	}

	return matches
}
