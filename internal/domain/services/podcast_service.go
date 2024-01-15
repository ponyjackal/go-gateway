package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ponyjackal/go-gateway/pkg/types"
	"github.com/spf13/viper"
)

type PodcastService struct {
	httpService *HTTPService
}

func NewPodcastService(httpService *HTTPService) *PodcastService {
	return &PodcastService{
		httpService: httpService,
	}
}

func (s *PodcastService) GetPodcasts(query types.GetPodcastsQuery) (*types.GetPodcastsResponse, error) {
	var PODCAST_SERVICE_URL = viper.GetString("PODCAST_SERVICE_URL")
	baseURL, err := url.Parse(fmt.Sprintf("%s/podcasts", PODCAST_SERVICE_URL))
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	// Prepare query parameters
	params := url.Values{}
	if query.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *query.Page))
	}
	if query.Limit != nil {
		params.Add("limit", fmt.Sprintf("%d", *query.Limit))
	}
	if query.Search != nil {
		params.Add("search", *query.Search)
	}
	if query.Title != nil {
		params.Add("title", *query.Title)
	}
	if query.CategoryName != nil {
		params.Add("categoryName", *query.CategoryName)
	}

	// Attach query parameters to the base URL
	baseURL.RawQuery = params.Encode()
	// Make the HTTP request
	res, err := s.httpService.Get(baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error making the request: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the JSON data
	var podcasts []types.Podcast
	err = json.Unmarshal(body, &podcasts)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return &types.GetPodcastsResponse{
		Podcasts: podcasts,
	}, nil
}
