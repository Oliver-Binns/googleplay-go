package authorization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticTokenSource_ReturnsTheProvidedToken(t *testing.T) {
	mockToken := "mock_token_value"
	tokenSource := StaticTokenSource(mockToken)

	token, err := tokenSource.Token()
	assert.Equal(t, token, mockToken)
	assert.Nil(t, err)
}
