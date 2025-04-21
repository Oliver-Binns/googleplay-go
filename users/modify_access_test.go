package users

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifyAccess_MakesRequest(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{ }`,
	}

	_, _ = ModifyAccess(
		httpClient, context.Background(), "https://example.com",
		"12345",
		[]AppLevelPermission{CanManageAppContent, CanManageDeeplinks},
	)

	assert.Equal(t, len(httpClient.requests), 1)
	assert.Equal(t, httpClient.requests[0].Method, "PATCH")
	assert.Equal(t, httpClient.requests[0].URL.String(), "https://example.com/12345?updateMask=appLevelPermissions")

	bodyBytes, err := io.ReadAll(httpClient.requests[0].Body)
	assert.NoError(t, err)
	bodyString := string(bodyBytes)
	assert.Equal(t,
		`{"appLevelPermissions":["CAN_MANAGE_APP_CONTENT","CAN_MANAGE_DEEPLINKS"]}
`, bodyString)
}

func TestModifyAccess_DecodesResponse(t *testing.T) {
	httpClient := &mockHTTPClient{
		response: `{
			"name": "45678",
			"appLevelPermissions": [
				"CAN_MANAGE_PUBLIC_LISTING","CAN_MANAGE_PUBLIC_APKS"
			]
		}`,
	}

	grant, err := ModifyAccess(
		httpClient, context.Background(), "https://example.com",
		"12345",
		[]AppLevelPermission{},
	)

	assert.NoError(t, err)
	assert.Equal(t, grant.Name, "45678")
	assert.Equal(t, grant.AppLevelPermissions, []AppLevelPermission{CanManagePublicListing, CanManagePublicAPKs})

}
