package auth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIKey_WithAPIKeyName(t *testing.T) {
	apiKeyName := "test-api-key-name"
	apiKeyPrivateKey := "test-api-private-key"

	apiKey, err := NewAPIKey(WithAPIKeyName(apiKeyName, apiKeyPrivateKey))
	assert.NoError(t, err)
	assert.Equal(t, apiKeyName, apiKey.Name)
	assert.Equal(t, apiKeyPrivateKey, apiKey.PrivateKey)
}

func TestNewAPIKey_WithLoadAPIKeyFromEnv(t *testing.T) {
	t.Setenv(nameEnvVar, "env-api-key-name")
	t.Setenv(privateKeyEnvVar, "env-api-private-key")

	apiKey, err := NewAPIKey(WithLoadAPIKeyFromEnv(true))
	assert.NoError(t, err)
	assert.Equal(t, "env-api-key-name", apiKey.Name)
	assert.Equal(t, "env-api-private-key", apiKey.PrivateKey)
}

func TestNewAPIKey_NoValidOption(t *testing.T) {
	_, err := NewAPIKey()
	assert.Error(t, err)
}

func TestLoadAPIKeyFromEnv(t *testing.T) {
	t.Setenv(nameEnvVar, "env-api-key-name")
	t.Setenv(privateKeyEnvVar, "env-api-private-key")

	apiKey, err := loadAPIKeyFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, "env-api-key-name", apiKey.Name)
	assert.Equal(t, "env-api-private-key", apiKey.PrivateKey)
}

func TestLoadAPIKeyFromEnv_MissingEnvVars(t *testing.T) {
	t.Setenv(nameEnvVar, "")
	t.Setenv(privateKeyEnvVar, "")

	_, err := loadAPIKeyFromEnv()
	assert.Error(t, err)
}

func TestLoadAPIKeyFromFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "api_key_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create the specific file .coinbase_cloud_api_key.json in the temporary directory
	keyFilePath := filepath.Join(tempDir, ".coinbase_cloud_api_key.json")
	apiKeyData := `{"name": "file-api-key-name", "privateKey": "file-api-private-key"}`
	err = os.WriteFile(keyFilePath, []byte(apiKeyData), 0600)
	assert.NoError(t, err)

	// Change working directory to the temp directory
	originalWd, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(originalWd)

	err = os.Chdir(tempDir)
	assert.NoError(t, err)

	// Load the API key from the file
	apiKey, err := loadAPIKeyFromFile()
	assert.NoError(t, err)
	assert.Equal(t, "file-api-key-name", apiKey.Name)
	assert.Equal(t, "file-api-private-key", apiKey.PrivateKey)
}
