package dto

type DiscographyRequest struct {
	BandName	string	`form:"band_name" validate:"required"`
}

type DiscographyResponse struct {
	Albums    []Album  `json:"items"`
}