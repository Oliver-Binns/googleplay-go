package authorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type tokenExchanger struct {
	httpClient  HTTPClient
	tokenSource TokenSource
	context     context.Context
}

func NewTokenExchanger(
	httpClient HTTPClient,
	tokenSource TokenSource,
	context context.Context,
) TokenSource {
	return &tokenExchanger{
		httpClient:  httpClient,
		tokenSource: tokenSource,
		context:     context,
	}
}

// tokenExchange handles the token exchange process with the OAuth2 server.
// It retrieves a new access token for the API using the service token JWT.
// It returns the access token string or an error if the exchange fails.
func (t *tokenExchanger) Token() (string, error) {
	token, err := t.tokenSource.Token()
	if err != nil {
		return "", err
	}

	url := "https://www.googleapis.com/oauth2/v4/token"
	body := fmt.Appendf(nil, `{
		"grant_type":    "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"assertion":     "%s",
	}`, token)

	req, err := http.NewRequestWithContext(t.context, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code from request: %d", resp.StatusCode)
	}

	tokenResp := new(tokenResponse)
	if err := json.NewDecoder(resp.Body).Decode(tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return "", fmt.Errorf("failed to close response body: %w", err)
	}

	return tokenResp.AccessToken, err
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}
