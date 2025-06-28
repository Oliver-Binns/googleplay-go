package networking

import (
	"fmt"
	"net/http"
)

type TokenSource interface {
	Token() (string, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type authorizedClient struct {
	httpClient  HTTPClient
	tokenSource TokenSource
}

func NewAuthorizedClient(httpClient HTTPClient, tokenSource TokenSource) HTTPClient {
	return &authorizedClient{
		httpClient:  httpClient,
		tokenSource: tokenSource,
	}
}

func (c *authorizedClient) Do(req *http.Request) (*http.Response, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return resp, nil
}
