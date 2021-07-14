package client

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	Get(string) ([]byte, int, error)
	Post(string, io.Reader) ([]byte, int, error)
}

type StandardClient struct {
	baseURL string

	httpClient *http.Client
}

var (
	_ Client = (*StandardClient)(nil)
)

func NewClient(baseURL string) *StandardClient {
	return &StandardClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *StandardClient) makeRequest(method string, url string, body io.Reader) ([]byte, int, error) {
	request, err := http.NewRequest(method, c.baseURL+url, body)
	if err != nil {
		return nil, -100, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, -101, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	return responseBody, response.StatusCode, nil
}

func (c *StandardClient) Get(url string) ([]byte, int, error) {
	return c.makeRequest("GET", url, nil)
}

func (c *StandardClient) Post(url string, data io.Reader) ([]byte, int, error) {
	return c.makeRequest("POST", url, data)
}