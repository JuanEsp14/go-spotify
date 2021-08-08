package dto

type Album struct {
	Name     string  `json:"name"`
	Released string  `json:"release_date"`
	Tracks   int     `json:"total_tracks"`
	Cover    []Cover `json:"images"`
}

type Cover struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}
