package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oliver-binns/googleplay-go/networking"
)

func GrantAccess(
	c networking.HTTPClient,
	ctx context.Context,
	rawURL string,
	name string,
	permissions []AppLevelPermission,
) (*Grant, error) {
	grant := Grant{
		Name:                name,
		PackageName:         name,
		AppLevelPermissions: permissions,
	}

	body := bytes.NewBuffer(nil)
	err := json.NewEncoder(body).Encode(grant)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	newGrant := new(Grant)
	if err := json.NewDecoder(resp.Body).Decode(newGrant); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	return newGrant, nil
}
