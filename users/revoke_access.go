package users

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/oliver-binns/googleplay-go/networking"
)

func RevokeAccess(
	c networking.HTTPClient,
	ctx context.Context,
	rawURL string,
	name string,
) error {
	// Parse the raw URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}
	// Add name to path: this is the package / app ID
	parsedURL.Path = path.Join(parsedURL.Path, name)

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, parsedURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}
