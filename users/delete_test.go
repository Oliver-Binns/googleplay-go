package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUser_MakesRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	_ = Delete(
		httpClient, context.Background(), "https://example.com",
		"john.doe@example.com",
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "DELETE")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com/john.doe@example.com")
	assert.Equal(t, httpClient.requests[0].Body, http.NoBody)
}

func TestDeleteUser_IsSuccessful(t *testing.T) {
	code := http.StatusNoContent

	httpClient := &mockHTTPClient{
		statusCode: &code,
		response:   `{ }`,
	}

	err := Delete(
		httpClient, context.Background(), "https://example.com",
		"john.doe@example.com",
	)

	assert.NoError(t, err)
}
