package pkg

import "github.com/JuanEsp14/go-spotify/pkg/dto"

type SpotifyAPIService interface {
	getArtists(responseType *dto.ArtistsResponse, band dto.DiscographyRequest, spotifyClient dto.SpotifyUser) error
	getDiscography(responseType *dto.DiscographyResponse, bandId string, spotifyClient dto.SpotifyUser) error
}
