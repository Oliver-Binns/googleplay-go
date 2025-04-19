package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/oliver-binns/googleplay-go/networking"
)

func Update(
	c networking.HTTPClient,
	ctx context.Context,
	rawURL string,
	email string,
	permissions *[]DeveloperLevelPermission,
) (*User, error) {
	// Parse the raw URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	parsedURL.Path = path.Join(parsedURL.Path, email)

	updateMask := []string{}

	user := UserUpdateRequest{}

	if permissions != nil {
		updateMask = append(updateMask, "developerAccountPermissions")
		user.DeveloperAccountPermissions = *permissions
	}

	body := bytes.NewBuffer(nil)
	err = json.NewEncoder(body).Encode(user)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Add the "updateMask" query parameter
	// This is a comma-separated list of fields to be updated
	query := parsedURL.Query()
	query.Set("updateMask", strings.Join(updateMask, ","))
	parsedURL.RawQuery = query.Encode()

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, parsedURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	newUser := new(User)
	if err := json.NewDecoder(resp.Body).Decode(newUser); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	return newUser, nil
}

type UserUpdateRequest struct {
	DeveloperAccountPermissions []DeveloperLevelPermission `json:"developerAccountPermissions,omitempty"`
}
