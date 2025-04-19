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
	client      *networking.HTTPClient
	developerID string
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
		client:      &client,
		developerID: developerID,
	}
}

func (c *Client) ListUsers(ctx context.Context) ([]users.User, error) {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	return users.List(*c.client, ctx, url)
}

func (c *Client) CreateUser(
	email string,
	permission []users.DeveloperLevelPermission,
	ctx context.Context,
) (*users.User, error) {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	newUserRequest := users.User{
		Email:                       email,
		DeveloperAccountPermissions: permission,
	}

	return users.Create(*c.client, ctx, url, newUserRequest)
}

func (c *Client) UpdateUser(
	email string,
	permissions *[]users.DeveloperLevelPermission,
	ctx context.Context,
) (*users.User, error) {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	return users.Update(*c.client, ctx, url, email, permissions)
}

func (c *Client) DeleteUser(email string, ctx context.Context) error {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	return users.Delete(*c.client, ctx, url, email)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
