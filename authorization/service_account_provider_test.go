package authorization

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenSource_Token(t *testing.T) {
	testStartTime := float64(time.Now().Unix())

	// Generate a token
	pk, b, err := generatePrivateKey()
	require.NoError(t, err)

	acc := mockServiceAccount(b)

	source, err := NewTokenSource(acc)
	require.NoError(t, err)

	token, err := source.Token()
	require.NoError(t, err)

	// Parse the token
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return pk.Public(), nil
	})
	require.NoError(t, err)

	// The correct key was used:
	// Signed with RSA
	assert.Equal(t, jwt.SigningMethodRS256.Alg(), parsed.Method.Alg())
	assert.Equal(t, acc.PrivateKeyID, parsed.Header["kid"])

	claims, ok := parsed.Claims.(jwt.MapClaims)
	require.True(t, ok)

	// Contains the correct issuer and subject:
	assert.Equal(t, acc.ClientEmail, claims["iss"])
	assert.Equal(t, acc.ClientEmail, claims["sub"])

	// Contains the correct audience
	assert.Equal(t, "https://firestore.googleapis.com/", claims["aud"])

	// Issued time is between the start of the test and now
	assert.GreaterOrEqual(t, testStartTime, claims["iat"])
	assert.LessOrEqual(t, float64(time.Now().Unix()), claims["iat"])

	// Expiration time is one hour after issued time
	assert.Equal(t, claims["iat"].(float64)+3600, claims["exp"])
}

func TestTokenSource_TokenRefresh(t *testing.T) {
	// GIVEN I have a valid token
	_, b, err := generatePrivateKey()
	require.NoError(t, err)

	acc := mockServiceAccount(b)
	source, err := ShortlivedTokenSource(acc)

	require.NoError(t, err)

	token, err := source.Token()
	require.NoError(t, err)

	same, err := source.Token()
	require.NoError(t, err)
	assert.Equal(t, token, same)

	// WHEN the token expires
	time.Sleep(time.Second)

	// THEN it should be refreshed: this currently fails!
	new, err := source.Token()
	require.NoError(t, err)
	assert.NotEqual(t, token, new)
}

func mockServiceAccount(b []byte) ServiceAccount {
	return ServiceAccount{
		Type:                    "service_account",
		ProjectID:               "elixir",
		PrivateKeyID:            "key-id",
		PrivateKey:              string(b),
		ClientEmail:             "google-play-console@test.iam.gserviceaccount.com",
		ClientID:                "issuer-id",
		AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
		TokenURI:                "https://oauth2.googleapis.com/token",
		AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
		ClientX509CertURL:       "https://www.googleapis.com/robot/v1/metadata/x509/google-play-console%40test.iam.gserviceaccount.com",
		UniverseDomain:          "googleapis.com",
	}
}

func generatePrivateKey() (*rsa.PrivateKey, []byte, error) {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	err = pk.Validate()
	keyBytes := x509.MarshalPKCS1PrivateKey(pk)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	})

	return pk, pemBytes, nil
}

func ShortlivedTokenSource(account ServiceAccount) (TokenSource, error) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(account.PrivateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &tokenSource{
		account:   account,
		pk:        pk,
		expiresIn: time.Nanosecond,
	}, nil
}
