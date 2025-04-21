package users

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrantAccess_MakesRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	_, _ = GrantAccess(
		httpClient, context.Background(), "https://example.com",
		"12345", []AppLevelPermission{CanManageDeeplinks},
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "POST")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t, bodyString,
		`{"name":"12345","packageName":"12345","appLevelPermissions":["CAN_MANAGE_DEEPLINKS"]}
`)
}

func TestGrantAccess_DecodesResponse(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"name": "98765",
			"appLevelPermissions": ["CAN_MANAGE_PERMISSIONS", "CAN_MANAGE_ORDERS"]
		}`,
	}

	grant, err := GrantAccess(
		httpClient, context.Background(), "https://example.com",
		"12345", []AppLevelPermission{CanManageDeeplinks},
	)

	assert.NoError(t, err)
	assert.Equal(t, grant.Name, "98765")
	assert.Equal(t, grant.AppLevelPermissions, []AppLevelPermission{CanManagePermissions, CanManageOrders})
}
