package users

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/oliver-binns/googleplay-go/networking"
)

func Delete(c networking.HTTPClient, ctx context.Context, baseURL string, email string) error {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}
	parsedURL.Path = path.Join(parsedURL.Path, email)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, parsedURL.String(), http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}
