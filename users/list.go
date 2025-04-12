package users

import (
	"context"
	"encoding/json"
	"fmt"
	"googleplay-go/networking"
	"net/http"
)

func ListUsers(c networking.HTTPClient, ctx context.Context, url string) ([]User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, _ := c.Do(req)

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
