package users

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_MakesRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	user := User{
		Name:                       "John Doe",
		Email:                      "john.doe@example.com",
		DeveloperAccountPermission: []DeveloperLevelPermission{"CAN_MANAGE_PERMISSIONS_GLOBAL"},
	}

	_, _ = Create(
		httpClient, context.Background(), "https://example.com",
		user,
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "POST")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t, bodyString,
		`{"name":"John Doe","email":"john.doe@example.com","developerAccountPermissions":["CAN_MANAGE_PERMISSIONS_GLOBAL"]}
`)
}

func TestCreateUser_DecodesResponse(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"name": "John Doe",
			"email": "john.doe@example.com",
			"accessState": "INVITED",
			"partial": false,
			"developerAccountPermissions": [
				"CAN_REPLY_TO_REVIEWS_GLOBAL"
			]
		}`,
	}

	user := User{}

	createdUser, err := Create(
		httpClient, context.Background(), "https://example.com",
		user,
	)

	assert.NoError(t, err)
	assert.Equal(t, createdUser.Name, "John Doe")
	assert.Equal(t, createdUser.Email, "john.doe@example.com")
	assert.Equal(t, createdUser.AccessState, Invited)
	assert.Equal(t, createdUser.Partial, false)
	assert.Equal(t, createdUser.DeveloperAccountPermission, []DeveloperLevelPermission{"CAN_REPLY_TO_REVIEWS_GLOBAL"})

}
