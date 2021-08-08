package pkg

import (
	"github.com/JuanEsp14/go-spotify/pkg/dto"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"net/http"
)

const (
	BaseUrl      = "https://api.spotify.com"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SpotifyAPIServiceImpl struct {
	client    HTTPClient
}

func NewSpotifyAPIServiceImpl(client HTTPClient) SpotifyAPIService{
	return &SpotifyAPIServiceImpl{client: client}
}

func (s *SpotifyAPIServiceImpl) getArtists(responseType *dto.ArtistsResponse, band dto.DiscographyRequest, spotifyClient dto.SpotifyUser)  error{
	request, err := getRequest(spotifyClient, "/v1/search")
	if err != nil {
		logrus.Errorf("Error generating request %s", err)
		return err
	}

	q := request.URL.Query()
	q.Add("query", band.BandName)
	q.Add("type", "artist")
	q.Add("limit", "5")
	request.URL.RawQuery = q.Encode()
	response, err := s.client.Do(request)
	defer response.Body.Close()

	if err != nil {
		logrus.Errorf("Error getting Token: %s", err)
		return err
	}

	if response.StatusCode != 200 {
		logrus.Errorf("Error getting Albums: %s", response.Status)
		return fmt.Errorf("Error getting Albums: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(responseType)
	if err != nil {
		return err
	}

	logrus.Info(responseType)
	return nil
}

func (s *SpotifyAPIServiceImpl)  getDiscography(responseType *dto.DiscographyResponse, bandId string, spotifyClient dto.SpotifyUser) error{
	uri := fmt.Sprintf("/v1/artists/%s/albums", bandId)
	request, err := getRequest(spotifyClient, uri)
	if err != nil {
		logrus.Errorf("Error generating request %s", err)
		return err
	}
	response, err := s.client.Do(request)
	defer response.Body.Close()

	if err != nil {
		logrus.Errorf("Error getting Albums: %s", err)
		return err
	}

	if response.StatusCode != 200 {
		logrus.Errorf("Error getting Albums: %s", response.Status)
		return fmt.Errorf("Error getting Albums: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(responseType)
	if err != nil {
		return err
	}

	logrus.Info(responseType)
	return nil
}


func getRequest(spotifyClient dto.SpotifyUser, uri string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", BaseUrl, uri)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("Error generating Request: %s", err)
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", spotifyClient.Token))
	return request, nil
}

/*		MOCK 	*/
type ClientMock struct { mock.Mock }

func (m *ClientMock) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	if len(m.ExpectedCalls) == 0 {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), nil
}
