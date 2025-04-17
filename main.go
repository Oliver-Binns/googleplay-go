package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/oliver-binns/googleplay-go/authorization"
	"github.com/oliver-binns/googleplay-go/networking"
	"github.com/oliver-binns/googleplay-go/users"
)

func main() {
	filename := os.Args[1]
	developer_id := os.Args[2]
	data, err := os.ReadFile(filename)
	check(err)
	rawJson := string(data)
	serviceAccount := authorization.ServiceAccount{}
	err = json.Unmarshal([]byte(rawJson), &serviceAccount)
	check(err)

	tokenSource, err := authorization.NewTokenSource(serviceAccount)
	check(err)
	tokenExchanger := authorization.NewTokenExchanger(http.DefaultClient, tokenSource, context.Background())
	client := networking.NewAuthorizedClient(http.DefaultClient, tokenExchanger)

	url := fmt.Sprintf(
		"https://androidpublisher.googleapis.com/androidpublisher/v3/developers/%s/users",
		developer_id,
	)

	usersList, err := users.List(client, context.Background(), url)
	check(err)
	fmt.Println("User list: ", usersList)

	newUserRequest := users.User{
		Name:  "John Doe",
		Email: "john.doe@oliverbinns.co.uk",
		DeveloperAccountPermission: []users.DeveloperLevelPermission{
			users.CanReplyToReviewsGlobal,
		},
	}

	newUser, err := users.Create(client, context.Background(), url, newUserRequest)
	check(err)
	fmt.Println("Created user: ", newUser)

	err = users.Delete(client, context.Background(), url, newUser.Email)
	check(err)
	fmt.Println("Deleted user: ", newUser.Email)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
