package client

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/data-collector/store"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
)

type Client interface {
	Get(string) ([]byte, int, error)
	Post(string, io.Reader) ([]byte, int, error)
}

type StandardClient struct {
	baseURL          string
	httpClient       *http.Client
	agentInformation types.AgentInformation
}

var (
	_ Client = (*StandardClient)(nil)
)

// NewClient returns an instanced HTTP client
func NewClient(baseURL string) (*StandardClient, error) {
	store := store.Instance()
	certDir := utils.GetVariable(consts.CERT_DIR)
	caCert, err := ioutil.ReadFile(certDir + "/" + utils.GetVariable(consts.CLIENT_CERT))
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &StandardClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		},
		agentInformation: store.GetAgentInformation(),
	}, nil
}

// makeRequest will make any HTTP request and also sends common data required for each request
func (c *StandardClient) makeRequest(method string, url string, body io.Reader) ([]byte, int, error) {
	request, err := http.NewRequest(method, c.baseURL+url, body)
	if err != nil {
		return nil, -100, err
	}

	request.Header.Add("Agent-ID", c.agentInformation.ID.String())
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, -101, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, -102, err
	}

	return responseBody, response.StatusCode, nil
}

// Get makes a GET request to a given URL
func (c *StandardClient) Get(url string) ([]byte, int, error) {
	return c.makeRequest("GET", url, nil)
}

// Post makes a POST request to a givn URL
func (c *StandardClient) Post(url string, data io.Reader) ([]byte, int, error) {
	return c.makeRequest("POST", url, data)
}
