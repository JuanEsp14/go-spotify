package dto

type ArtistsResponse struct {
	Artists		ArtistsSearch `json:"artists"`
}

type ArtistsSearch struct {
	Items	 []Artist 	 `json:"items"`
	Total    int         `json:"total"`
}

type Artist struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	URI  string `json:"uri"`
	Endpoint     string            `json:"href"`
	ExternalURLs map[string]string `json:"external_urls"`
}

