package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevokeAccess_MakesRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	_ = RevokeAccess(
		httpClient, context.Background(), "https://example.com",
		"12345",
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "DELETE")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com/12345")
	assert.Equal(t, httpClient.requests[0].Body, nil)
}

func TestRevokeAccess_IsSuccessful(t *testing.T) {
	code := http.StatusNoContent

	httpClient := &mockHTTPClient{
		statusCode: &code,
		response:   `{ }`,
	}

	err := Delete(
		httpClient, context.Background(), "https://example.com",
		"98765",
	)

	assert.NoError(t, err)
}
