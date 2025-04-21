package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/oliver-binns/googleplay-go/networking"
)

func ModifyAccess(
	c networking.HTTPClient,
	ctx context.Context,
	rawURL string,
	name string,
	permissions []AppLevelPermission,
) (*Grant, error) {
	// Parse the raw URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	// Add name to path: this is the package / app ID
	parsedURL.Path = path.Join(parsedURL.Path, name)

	// Add the "updateMask" query parameter
	// This is a comma-separated list of fields to be updated
	query := parsedURL.Query()
	query.Set("updateMask", "appLevelPermissions")
	parsedURL.RawQuery = query.Encode()

	grant := Grant{
		AppLevelPermissions: permissions,
	}
	body := bytes.NewBuffer(nil)
	err = json.NewEncoder(body).Encode(grant)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, parsedURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	// Decode the response
	newGrant := new(Grant)
	if err := json.NewDecoder(resp.Body).Decode(newGrant); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	return newGrant, nil
}
