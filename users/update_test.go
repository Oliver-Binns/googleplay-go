package users

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUsers_MakesRequestWithNoValues(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	_, _ = Update(
		httpClient, context.Background(), "https://example.com",
		nil,
		nil,
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "PATCH")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com?updateMask=")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t, bodyString,
		`{}
`)
}

func TestUpdateUsers_MakesRequestWithName(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	name := "John Doe"
	_, _ = Update(
		httpClient, context.Background(), "https://example.com",
		&name,
		nil,
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "PATCH")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com?updateMask=name")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t, bodyString,
		`{"name":"John Doe"}
`)
}

func TestUpdateUsers_MakesRequestWithPermissions(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	permissions := []DeveloperLevelPermission{
		CanManagePermissionsGlobal,
		CanReplyToReviewsGlobal,
	}
	_, _ = Update(
		httpClient, context.Background(), "https://example.com",
		nil,
		&permissions,
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "PATCH")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com?updateMask=developerAccountPermissions")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t, bodyString,
		`{"developerAccountPermissions":["CAN_MANAGE_PERMISSIONS_GLOBAL","CAN_REPLY_TO_REVIEWS_GLOBAL"]}
`)
}

func TestUpdateUsers_MakesRequestWithAllParameters(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	name := "John Doe"
	permissions := []DeveloperLevelPermission{
		CanManagePermissionsGlobal,
		CanReplyToReviewsGlobal,
	}
	_, _ = Update(
		httpClient, context.Background(), "https://example.com",
		&name,
		&permissions,
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "PATCH")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com?updateMask=name%2CdeveloperAccountPermissions")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t, bodyString,
		`{"name":"John Doe","developerAccountPermissions":["CAN_MANAGE_PERMISSIONS_GLOBAL","CAN_REPLY_TO_REVIEWS_GLOBAL"]}
`)
}

func TestUpdateUsers_DecodesResponse(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"name": "Oliver Binns",
			"email": "mail@oliverbinns.co.uk",
			"accessState": "ACCESS_GRANTED",
			"partial": false,
			"developerAccountPermissions": [ ],
			"grants": [ ]
		}`,
	}

	user, _ := Update(
		httpClient, context.Background(), "https://example.com",
		nil,
		nil,
	)

	assert.Equal(t, user.Name, "Oliver Binns")
	assert.Equal(t, user.Email, "mail@oliverbinns.co.uk")
	assert.Equal(t, user.AccessState, AccessGranted)
}
