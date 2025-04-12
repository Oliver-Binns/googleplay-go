package googleplay

import (
	"googleplay-go/authorization"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient  HTTPClient
	tokenSource authorization.TokenSource
}

func NewClient(httpClient HTTPClient, tokenSource authorization.TokenSource) *Client {
	return &Client{
		httpClient:  httpClient,
		tokenSource: tokenSource,
	}
}
