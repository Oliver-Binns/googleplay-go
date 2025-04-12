package users

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsers_MakesRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"users": [
			
			],
			"nextPageToken": string
		}`,
	}

	_, _ = ListUsers(
		httpClient, context.Background(), "https://example.com",
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "GET")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com")
}

type mockHTTPClient struct {
	requests []*http.Request
	response string
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.requests = append(c.requests, req)

	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(c.response)))

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       responseBody,
	}, nil
}
