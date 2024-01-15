package services

import (
	"io"
	"net/http"
	"time"
)

type HTTPService struct {
	client *http.Client
}

func NewHTTPService() *HTTPService {
	return &HTTPService{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *HTTPService) Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return s.client.Do(req)
}

func (s *HTTPService) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return s.client.Do(req)
}

func (s *HTTPService) Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return s.client.Do(req)
}

func (s *HTTPService) Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return s.client.Do(req)
}

func (s *HTTPService) Delete(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return s.client.Do(req)
}
