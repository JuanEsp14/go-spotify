package dto

type SpotifyUser struct {
	ClientId 		string `form:"client_id" validate:"required"`
	ClientSecret	string `form:"client_secret" validate:"required"`
	Token 			string `form:"token" validate:"required"`
}