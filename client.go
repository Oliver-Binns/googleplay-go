package googleplay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/oliver-binns/googleplay-go/authorization"
	"github.com/oliver-binns/googleplay-go/networking"
	"github.com/oliver-binns/googleplay-go/users"
)

type Client struct {
	client  *networking.HTTPClient
	baseURL string
}

func GooglePlayClient(developerID string, serviceAccountJson string) *Client {
	serviceAccount := authorization.ServiceAccount{}
	err := json.Unmarshal([]byte(serviceAccountJson), &serviceAccount)
	check(err)

	tokenSource, err := authorization.NewTokenSource(serviceAccount)
	check(err)

	tokenExchanger := authorization.NewTokenExchanger(http.DefaultClient, tokenSource, context.Background())
	client := networking.NewAuthorizedClient(http.DefaultClient, tokenExchanger)

	return googlePlayClient(developerID, client)
}

func GooglePlayClientWithToken(developerID string, accessToken string) *Client {
	tokenSource := authorization.StaticTokenSource(accessToken)
	client := networking.NewAuthorizedClient(http.DefaultClient, tokenSource)

	return googlePlayClient(developerID, client)
}

func googlePlayClient(developerID string, client networking.HTTPClient) *Client {
	return &Client{
		client: &client,
		baseURL: fmt.Sprintf(
			"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
			developerID,
		),
	}
}

func (c *Client) ListUsers(ctx context.Context) ([]users.User, error) {
	return users.List(*c.client, ctx, c.baseURL)
}

func (c *Client) CreateUser(
	email string,
	permission []users.DeveloperLevelPermission,
	ctx context.Context,
) (*users.User, error) {
	newUserRequest := users.User{
		Email:                       email,
		DeveloperAccountPermissions: permission,
	}

	return users.Create(*c.client, ctx, c.baseURL, newUserRequest)
}

func (c *Client) UpdateUser(
	email string,
	permissions *[]users.DeveloperLevelPermission,
	ctx context.Context,
) (*users.User, error) {
	return users.Update(*c.client, ctx, c.baseURL, email, permissions)
}

func (c *Client) DeleteUser(email string, ctx context.Context) error {
	return users.Delete(*c.client, ctx, c.baseURL, email)
}

func (c *Client) GrantAccess(
	email string,
	appID string,
	permissions []users.AppLevelPermission,
	ctx context.Context,
) (*users.Grant, error) {
	url, err := c.addToPath([]string{email, "grants"})

	if err != nil {
		return nil, err
	}

	return users.GrantAccess(*c.client, ctx, *url, appID, permissions)
}

func (c *Client) ModifyAccess(
	email string,
	appID string,
	permissions []users.AppLevelPermission,
	ctx context.Context,
) (*users.Grant, error) {
	url, err := c.addToPath([]string{email, "grants"})

	if err != nil {
		return nil, err
	}

	return users.ModifyAccess(*c.client, ctx, *url, appID, permissions)
}

func (c *Client) RevokeAccess(
	email string,
	appID string,
	ctx context.Context,
) error {
	url, err := c.addToPath([]string{email, "grants"})

	if err != nil {
		return err
	}

	return users.RevokeAccess(*c.client, ctx, *url, appID)
}

func (c *Client) addToPath(components []string) (*string, error) {
	// Parse the raw URL
	parsedURL, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	// Add name to path: this is the package / app ID
	for _, component := range components {
		parsedURL.Path = path.Join(parsedURL.Path, component)
	}

	completeURL := parsedURL.String()

	return &completeURL, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
