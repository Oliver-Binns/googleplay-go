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

func (c *Client) ListUsers() ([]users.User, error) {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	usersList, err := users.List(*c.client, context.Background(), url)
	if err != nil {
		return nil, err
	}
	return usersList, nil
}

func (c *Client) CreateUser(
	name string,
	email string,
	permission []users.DeveloperLevelPermission,
) (*users.User, error) {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	newUserRequest := users.User{
		Name:                       name,
		Email:                      email,
		DeveloperAccountPermission: permission,
	}

	newUser, err := users.Create(*c.client, context.Background(), url, newUserRequest)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (c *Client) DeleteUser(email string) error {
	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		c.developerID,
	)

	err := users.Delete(*c.client, context.Background(), url, email)
	if err != nil {
		return err
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
