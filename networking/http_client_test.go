package networking

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthorizedClient_AuthorizationHeaderIsSet(t *testing.T) {
	httpClient := &mockHTTPClient{
		statusCode: http.StatusOK,
	}
	tokenSource := &mockTokenSource{}

	authorizedClient := NewAuthorizedClient(httpClient, tokenSource)
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	_, err := authorizedClient.Do(req)

	assert.Nil(t, err, "Unexpected error: %v", err)
	assert.Equal(
		t, "Bearer test-token", req.Header.Get("Authorization"),
		"Authorization header not set correctly",
	)
}

func TestNewAuthorizedClient_Allows2XXStatusCodes(t *testing.T) {
	httpClient := &mockHTTPClient{
		statusCode: http.StatusNoContent,
	}
	tokenSource := &mockTokenSource{}

	authorizedClient := NewAuthorizedClient(httpClient, tokenSource)
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	_, err := authorizedClient.Do(req)

	assert.NoError(t, err)
}

type mockTokenSource struct{}

func (c *mockTokenSource) Token() (string, error) {
	return "test-token", nil
}

type mockHTTPClient struct {
	statusCode int
	requests   []*http.Request
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.requests = append(c.requests, req)

	return &http.Response{
		StatusCode: c.statusCode,
	}, nil
}
