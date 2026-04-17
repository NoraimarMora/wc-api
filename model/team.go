package model

type Team struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Confederation string `json:"confederation"`
	FlagURl       string `json:"flag_uri"`
	Group         string `json:"group"`
	WorldRanking  int    `json:"world_ranking"`
	Appearances   int    `json:"appearances"`
	Host          bool   `json:"host"`
	Colors        Colors `json:"colors"`
}

type Colors struct {
	PrimaryColor       string `json:"primary_color"`
	SecondaryColor     string `json:"secondary_color"`
	PrimaryTextColor   string `json:"primary_text_color"`
	SecondaryTextColor string `json:"secondary_text_color"`
}

type Teams map[string]*Team

func (t Teams) ToSlice() []*Team {
	slice := make([]*Team, 0, len(t))
	for _, v := range t {
		slice = append(slice, v)
	}

	return slice
}

func (t Teams) ByGroup(group string) Teams {
	if group == "" {
		return t
	}

	teams := make(Teams)
	for k, v := range t {
		if v.Group == group {
			teams[k] = v
		}
	}

	return teams
}

func (t Teams) ByConfederation(confederation string) Teams {
	if confederation == "" {
		return t
	}

	teams := make(Teams)
	for k, v := range t {
		if v.Confederation == confederation {
			teams[k] = v
		}
	}

	return teams
}
