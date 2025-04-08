package authorization

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ensure that ServiceAccount can be decoded from JSON
func TestJsonDecoding(t *testing.T) {
	rawJson := `{
		"type": "service_account",
		"project_id": "xyz",
		"private_key_id": "key-id",
		"private_key": "-----BEGIN PRIVATE KEY-----MIIEvQ=-----END PRIVATE KEY-----",
		"client_email": "google-play-console@test.iam.gserviceaccount.com",
		"client_id": "000000000000000000000",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/google-play-console%40test.iam.gserviceaccount.com",
		"universe_domain": "googleapis.com"
	}`

	account := ServiceAccount{}
	err := json.Unmarshal([]byte(rawJson), &account)
	assert.NoError(t, err)

	assert.Equal(t, account.Type, "service_account")
	assert.Equal(t, account.ProjectID, "xyz")
	assert.Equal(t, account.PrivateKeyID, "key-id")
	assert.Equal(t, account.PrivateKey, "-----BEGIN PRIVATE KEY-----MIIEvQ=-----END PRIVATE KEY-----")
	assert.Equal(t, account.ClientEmail, "google-play-console@test.iam.gserviceaccount.com")
	assert.Equal(t, account.ClientID, "000000000000000000000")
	assert.Equal(t, account.AuthURI, "https://accounts.google.com/o/oauth2/auth")
	assert.Equal(t, account.TokenURI, "https://oauth2.googleapis.com/token")
	assert.Equal(t, account.AuthProviderX509CertURL, "https://www.googleapis.com/oauth2/v1/certs")
	assert.Equal(t, account.ClientX509CertURL, "https://www.googleapis.com/robot/v1/metadata/x509/google-play-console%40test.iam.gserviceaccount.com")
	assert.Equal(t, account.UniverseDomain, "googleapis.com")
}
