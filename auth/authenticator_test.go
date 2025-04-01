package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2/jwt"
)

func generateTestPrivateKey() (string, *ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", nil, err
	}

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return "", nil, err
	}

	pemKey := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	})

	return string(pemKey), key, nil
}

func TestBuildJWT(t *testing.T) {
	privateKey, ecdsaPrivateKey, err := generateTestPrivateKey()
	assert.NoError(t, err)

	apiKey := &APIKey{
		Name:       "test-api-key",
		PrivateKey: privateKey,
	}

	authenticator := NewAuthenticator(apiKey)
	service := "test-service"
	uri := "test-uri"

	jwtString, err := authenticator.BuildJWT(service, uri)
	assert.NoError(t, err)
	assert.NotEmpty(t, jwtString)

	parsedJWT, err := jwt.ParseSigned(jwtString)
	assert.NoError(t, err)

	claims := &APIKeyClaims{}
	err = parsedJWT.Claims(&ecdsaPrivateKey.PublicKey, claims)
	assert.NoError(t, err)
	assert.Equal(t, apiKey.Name, claims.Subject)
	assert.Equal(t, "coinbase-cloud", claims.Issuer)
	assert.WithinDuration(t, time.Now(), claims.NotBefore.Time(), time.Minute)
	assert.WithinDuration(t, time.Now().Add(1*time.Minute), claims.Expiry.Time(), time.Minute)
	assert.Equal(t, service, claims.Audience[0])
	assert.Equal(t, uri, claims.URI)
}
