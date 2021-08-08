package pkg

import (
	"github.com/JuanEsp14/go-spotify/pkg/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)


type SpotifyGinController struct {
	spotifyAPIService SpotifyAPIService
	spotifyUser		  dto.SpotifyUser
}

func NewSpotifyGinController(spotifyAPIService SpotifyAPIService) *SpotifyGinController {
	return &SpotifyGinController{spotifyAPIService: spotifyAPIService}
}

func (m *SpotifyGinController) SetupRouter(server *gin.Engine) {
	server.POST("/setToken", m.SetToken)
	server.GET("/getDiscography", m.GetDiscography)
}

// setToken godoc
// @tags Spotify
// @Summary post
// @Description Introduce a name of the band
// @Accept  json
// @Produce  json
// @Router /setToken [post]
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Param client_id query string true "String ClientId: client account id"
// @Param client_secret query string false "String ClientSecret: client secret to spotify"
// @Param token query string false "String Token: token to use in request"
func (m *SpotifyGinController) SetToken(context *gin.Context){
	logrus.Info("Init set Token")
	if err := context.BindQuery(&m.spotifyUser); err != nil {
		context.JSON(http.StatusBadRequest, "Error getting Query params")
		return
	}
	if err := validator.New().Struct(m.spotifyUser); err != nil {
		logrus.Error(fmt.Sprintf("Input validation FAILED. Error: %s", err))
		context.JSON(http.StatusBadRequest, fmt.Sprintf("Error input data: %s", err))
		return
	}

	logrus.Info(fmt.Sprintf("User setting with clientId: %s", m.spotifyUser.ClientId))
	context.JSON(http.StatusOK, "User setting")
}

// getDiscography godoc
// @tags Spotify
// @Summary get
// @Description Introduce a name of the band
// @Accept  json
// @Produce  json
// @Router /getDiscography [get]
// @Success 200 {object} []dto.Album
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Param band_name query string true "String Band Name: name of the band"
func (m *SpotifyGinController)  GetDiscography(context *gin.Context){
	request, err := m.validateData(context)
	if err != nil {
		logrus.Error("Error validating data")
		return
	}
	artists := new(dto.ArtistsResponse)
	err = m.spotifyAPIService.getArtists(artists, request, m.spotifyUser)

	if err != nil {
		logrus.Errorf("Error getting Artists: %s", err)
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	discography := new(dto.DiscographyResponse)
	err = m.spotifyAPIService.getDiscography(discography, artists.Artists.Items[0].ID, m.spotifyUser)

	if err != nil {
		logrus.Errorf("Error getting Discography: %s", err)
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, discography.Albums)
}

func (m *SpotifyGinController) validateData(context *gin.Context) (dto.DiscographyRequest, error) {
	if err := validator.New().Struct(m.spotifyUser); err != nil {
		logrus.Error(fmt.Sprintf("Input validation FAILED. Error: %s", err))
		context.JSON(http.StatusBadRequest, "You must enter your credentials with setToken endpoint to use the API")
		return dto.DiscographyRequest{}, err
	}
	request := dto.DiscographyRequest{}
	if err := context.BindQuery(&request); err != nil {
		context.JSON(http.StatusBadRequest, "Error getting Query params")
		return dto.DiscographyRequest{}, err
	}
	if err := validator.New().Struct(request); err != nil {
		logrus.Error(fmt.Sprintf("Input validation FAILED. Error: %s", err))
		context.JSON(http.StatusBadRequest, fmt.Sprintf("Input validation FAILED. Error: %s", err))
		return dto.DiscographyRequest{}, err
	}
	return request, nil
}