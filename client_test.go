package googleplay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToPath_ValidURL(t *testing.T) {
	client := &Client{
		baseURL: "https://androidpublisher.googleapis.com/androidpublisher/v3/developers/123/users",
	}

	expected := "https://androidpublisher.googleapis.com/androidpublisher/v3/developers/123/users/456/grants"
	actual, err := client.addToPath([]string{"456", "grants"})

	assert.Nil(t, err)
	assert.Equal(t, expected, *actual)
}
