package authorization

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenExchanger_MakesTokenExchangeRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		statusCode: http.StatusOK,
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

	request := httpClient.requests[0]
	assert.Equal(t, request.Method, "POST")
	assert.Equal(t, request.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, request.URL.String(), "https://www.googleapis.com/oauth2/v4/token")

	tokenExchangeRequestBody := new(tokenExchangeRequest)
	err := json.NewDecoder(request.Body).Decode(tokenExchangeRequestBody)
	assert.Nil(t, err)

	assert.Equal(t, tokenExchangeRequestBody.GrantType, "urn:ietf:params:oauth:grant-type:jwt-bearer")
	assert.Equal(t, tokenExchangeRequestBody.Assertion, "test-token")
}

func TestTokenExchanger_ReturnsAccessToken(t *testing.T) {
	httpClient := &mockHTTPClient{
		statusCode: http.StatusOK,
		response: `{
			"access_token": "test-access-token"
		}`,
	}

	tokenExchanger := NewTokenExchanger(
		httpClient,
		&mockTokenSource{},
		context.Background(),
	)

	accessToken, err := tokenExchanger.Token()
	assert.Equal(t, accessToken, "test-access-token")
	assert.Nil(t, err)
}

func TestTokenExchanger_ReturnsErrorForNon200HTTPStatus(t *testing.T) {
	httpClient := &mockHTTPClient{
		statusCode: http.StatusUnauthorized,
		response: `{
			"access_token": "test-access-token"
		}`,
	}

	tokenExchanger := NewTokenExchanger(
		httpClient,
		&mockTokenSource{},
		context.Background(),
	)

	accessToken, err := tokenExchanger.Token()
	assert.Equal(t, err.Error(), "unexpected status code from request: 401")
	assert.Equal(t, accessToken, "")
}

func TestTokenExchanger_ReturnsErrorWhenUnableToDecode(t *testing.T) {
	httpClient := &mockHTTPClient{
		statusCode: http.StatusOK,
		response:   ``,
	}

	tokenExchanger := NewTokenExchanger(
		httpClient,
		&mockTokenSource{},
		context.Background(),
	)

	accessToken, err := tokenExchanger.Token()
	assert.Equal(t, err.Error(), "failed to decode response: EOF")
	assert.Equal(t, accessToken, "")
}

type mockTokenSource struct{}

func (c *mockTokenSource) Token() (string, error) {
	return "test-token", nil
}

type mockHTTPClient struct {
	requests   []*http.Request
	statusCode int
	response   string
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.requests = append(c.requests, req)

	responseBody := io.NopCloser(bytes.NewReader([]byte(c.response)))

	return &http.Response{
		StatusCode: c.statusCode,
		Body:       responseBody,
	}, nil
}
