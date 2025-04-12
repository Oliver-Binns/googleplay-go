package main

import (
	"context"
	"encoding/json"
	"fmt"
	"googleplay-go/authorization"
	"googleplay-go/networking"
	"googleplay-go/users"
	"net/http"
	"os"
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

	users, err := users.ListUsers(client, context.Background(), url)
	check(err)

	fmt.Println("Hello World", users)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
