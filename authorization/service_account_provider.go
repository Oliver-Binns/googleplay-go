package authorization

import (
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenSource interface {
	Token() (string, error)
}

func NewTokenSource(account ServiceAccount) (TokenSource, error) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(account.PrivateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &tokenSource{
		account:   account,
		pk:        pk,
		expiresIn: time.Hour,
	}, nil
}

type tokenSource struct {
	sync.Mutex

	account   ServiceAccount
	pk        *rsa.PrivateKey
	expiresIn time.Duration
	bearer    string
	expireAt  time.Time
}

func (ts *tokenSource) Token() (string, error) {
	ts.Lock()
	defer ts.Unlock()

	if ts.isExpired() {
		return ts.refresh()
	}

	return ts.bearer, nil
}

func (ts *tokenSource) isExpired() bool {
	return time.Now().After(ts.expireAt)
}

func (ts *tokenSource) refresh() (string, error) {
	// Create JWT as defined in https://developers.google.com/identity/protocols/oauth2/service-account
	iat := time.Now()
	exp := iat.Add(ts.expiresIn)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":   ts.account.ClientEmail,
		"sub":   ts.account.ClientEmail,
		"scope": "https://www.googleapis.com/auth/androidpublisher",
		"aud":   "https://oauth2.googleapis.com/token",
		"iat":   iat.Unix(),
		"exp":   exp.Unix(),
	})
	token.Header["alg"] = "RS256"
	token.Header["typ"] = "JWT"
	token.Header["kid"] = ts.account.PrivateKeyID

	bearer, err := token.SignedString(ts.pk)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	ts.bearer = bearer
	ts.expireAt = exp

	return bearer, nil
}
