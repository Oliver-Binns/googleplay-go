package googleplay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
