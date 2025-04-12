package users

import (
	"bytes"
	"context"
	"io"
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

func TestListUsers_DecodesResponse(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"users": [
				{
					"name": "Oliver Binns",
					"email": "mail@oliverbinns.co.uk",
					"accessState": "ACCESS_GRANTED",
					"partial": false,
					"developerAccountPermissions": [ ],
					"grants": [ ]
				}
			],
			"nextPageToken": ""
		}`,
	}

	users, _ := ListUsers(
		httpClient, context.Background(), "https://example.com",
	)

	assert.Equal(t, len(users), 1)

	assert.Equal(t, users[0].Name, "Oliver Binns")
	assert.Equal(t, users[0].Email, "mail@oliverbinns.co.uk")
	assert.Equal(t, users[0].AccessState, AccessGranted)
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
