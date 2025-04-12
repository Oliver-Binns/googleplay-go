package authorization

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenExchanger_MakesTokenExchangeRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"access_token": "test-access-token"
		}`,
	}

	tokenExchanger := NewTokenExchanger(
		httpClient,
		&mockTokenSource{},
		context.Background(),
	)

	_, _ = tokenExchanger.Token()

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "POST")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://www.googleapis.com/oauth2/v4/token")
}

type mockTokenSource struct{}

func (c *mockTokenSource) Token() (string, error) {
	return "test-token", nil
}

type mockHTTPClient struct {
	requests []*http.Request
	response string
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.requests = append(c.requests, req)

	responseBody := io.NopCloser(bytes.NewReader([]byte(c.response)))

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       responseBody,
	}, nil
}
