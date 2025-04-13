package users

import (
	"context"
	"encoding/json"
	"fmt"
	"googleplay-go/networking"
	"net/http"
	"net/url"
)

func List(c networking.HTTPClient, ctx context.Context, rawURL string) ([]User, error) {
	// Parse the raw URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	// Add the "page_size" query parameter
	query := parsedURL.Query()
	query.Set("page_size", "-1")
	parsedURL.RawQuery = query.Encode()

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	userListResponse := new(userListResponse)
	if err := json.NewDecoder(resp.Body).Decode(userListResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	return userListResponse.Users, nil
}

type userListResponse struct {
	Users         []User `json:"users"`
	NextPageToken string `json:"nextPageToken"`
}
