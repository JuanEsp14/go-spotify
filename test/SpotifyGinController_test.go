package test

import (
	"bytes"
	"github.com/JuanEsp14/go-spotify/pkg"
	"github.com/JuanEsp14/go-spotify/pkg/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	BaseUrl     	= "https://api.spotify.com"
	BandName		= "mock-band-name"
	BandName2		= "mock-band-name-2"
	BandId			= "22bE4uQ6baNwSHPVcDxLCe"
	BandId2			= "6n5ceqTPlycEbZZLpGr3Tb"
	ClientId		= "mock-client-id"
	ClientSecret	= "mock-client-secret"
	Token			= "mock-token"
)

var controller *pkg.SpotifyGinController
var router *gin.Engine
var mockClient pkg.ClientMock

var _ = BeforeSuite(setupTests)

func TestSpotifyGinController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spotify controller Suite")
}

func setupTests() {
	mockClient = pkg.ClientMock{}
	serviceSpotify := pkg.NewSpotifyAPIServiceImpl(&mockClient)
	controller = pkg.NewSpotifyGinController(serviceSpotify)
	router = gin.Default()
	controller.SetupRouter(router)
}

var _ = Describe("Spotify Gin Controller", func() {
	Context("Return error getting discography without spotifyClient", getDiscographyWithoutUser)
	Context("Return error validating data", setUserWithoutClientId)
	Context("Set a Spotify user in controller", setUserHappyPath)
	Context("Return error getting discography without band name", getDiscographyWithoutBandName)
	Context("Return error getting discography without bad token", getDiscographyUnauthorizedOnArtists)
	Context("Return error getting discography without bad token", getDiscographyUnauthorizedOnDiscography)
	Context("Return discography of the band", getDiscographyHappyPath)
})


func getDiscographyWithoutUser() {
	When("Get discography without SpotifyClient", func() {
		It("Should get http status 400 with a message requesting data", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getDiscography", nil)

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusBadRequest))
			Ω(w.Body.String()).To(Equal(`"You must enter your credentials with setToken endpoint to use the API"`))
		})
	})
}

func setUserWithoutClientId() {
	When("Send a request without client Id", func() {
		It("Should get http status 400 with a message validating data", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/setToken", nil)
			q := req.URL.Query()
			q.Add("client_secret", ClientSecret)
			q.Add("token", Token)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusBadRequest))
			Ω(w.Body.String()).To(Equal(`"Error input data: Key: 'SpotifyUser.ClientId' Error:Field validation for 'ClientId' failed on the 'required' tag"`))
		})
	})
}

func setUserHappyPath() {
	When("Send a clientId, clientSecret and token", func() {
		It("Should get http status 200 with message User setting", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/setToken", nil)
			q := req.URL.Query()
			q.Add("client_id", ClientId)
			q.Add("client_secret", ClientSecret)
			q.Add("token", Token)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusOK))
			Ω(w.Body.String()).To(Equal(`"User setting"`))
		})
	})
}

func getDiscographyWithoutBandName() {
	When("Get discography without band name", func() {
		It("Should get http status 400 with a message validating data", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getDiscography", nil)

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusBadRequest))
			Ω(w.Body.String()).To(Equal(`"Input validation FAILED. Error: Key: 'DiscographyRequest.BandName' Error:Field validation for 'BandName' failed on the 'required' tag"`))
		})
	})
}

func getDiscographyUnauthorizedOnArtists() {
	When("Get discography with band name and bad token", func() {
		It("Should get http status 400 with message", func() {
			artistRequest := getArtistRequest(BandName)
			artistResponse := getArtistResponse("artists_response.json")
			artistResponse.StatusCode = 401
			mockClient.On("Do", artistRequest).Return(artistResponse)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getDiscography", nil)
			q := req.URL.Query()
			q.Add("band_name", BandName)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusBadRequest))
		})
	})
}

func getDiscographyUnauthorizedOnDiscography() {
	When("Get discography with band name and bad token", func() {
		It("Should get http status 400 with message", func() {
			artistRequest := getArtistRequest(BandName)
			artistResponse := getArtistResponse("artists_response.json")
			artistResponse.StatusCode = 200
			discographyRequest := getDiscographyRequest(BandId)
			discographyResponse := http.Response{}
			discographyResponse.StatusCode = 401
			mockClient.On("Do", artistRequest).Return(artistResponse)
			mockClient.On("Do", discographyRequest).Return(discographyResponse)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getDiscography", nil)
			q := req.URL.Query()
			q.Add("band_name", BandName)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusBadRequest))
		})
	})
}


func getDiscographyHappyPath() {
	When("Get discography with band name", func() {
		It("Should get http status 200 with mock discography", func() {
			artistRequest := getArtistRequest(BandName2)
			artistResponse := getArtistResponse("artists_response_2.json")
			artistResponse.StatusCode = 200
			discographyRequest := getDiscographyRequest(BandId2)
			discographyResponse := getDiscographyResponse()
			discographyResponse.StatusCode = 200
			mockClient.On("Do", artistRequest).Return(artistResponse)
			mockClient.On("Do", discographyRequest).Return(discographyResponse)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getDiscography", nil)
			q := req.URL.Query()
			q.Add("band_name", BandName2)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)
			Ω(w.Code).To(Equal(http.StatusOK))
		})
	})
}

func getArtistRequest(bandName string) *http.Request {
	spotifyClient := dto.SpotifyUser{Token: Token}
	request, _ := getRequest(spotifyClient, "/v1/search")
	q := request.URL.Query()
	q.Add("query", bandName)
	q.Add("type", "artist")
	q.Add("limit", "5")
	request.URL.RawQuery = q.Encode()
	return request
}

func getArtistResponse(jsonResponse string) *http.Response {
	jsonFile, _ := os.Open(jsonResponse)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(byteValue))}
	return &response
}

func getDiscographyRequest(bandId string) *http.Request {
	spotifyClient := dto.SpotifyUser{Token: Token}
	request, _ := getRequest(spotifyClient, fmt.Sprintf("/v1/artists/%s/albums", bandId))
	return request
}

func getDiscographyResponse() *http.Response {
	jsonFile, _ := os.Open("discography_response.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(byteValue))}
	return &response
}


func getRequest(spotifyClient dto.SpotifyUser, uri string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", BaseUrl, uri)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", spotifyClient.Token))
	return request, nil
}
