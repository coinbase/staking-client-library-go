package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const (
	// NameEnvVar is the environment variable for the Coinbase Cloud API Key name.
	nameEnvVar = "COINBASE_CLOUD_API_KEY_NAME"

	// PrivateKeyEnvVar is the environment variable for the Coinbase Cloud API Key private key.
	privateKeyEnvVar = "COINBASE_CLOUD_API_KEY_PRIVATE_KEY"
)

// APIKey represents a Coinbase Cloud API Key.
type APIKey struct {
	Name       string `json:"name"`
	PrivateKey string `json:"privateKey"`
}

// apiKeyConfig contains configuration information for constructing the API Key.
type apiKeyConfig struct {
	loadAPIKeyFromEnv  bool
	loadAPIKeyFromFile bool
	apiKeyName         string
	apiKeyPrivateKey   string
}

// APIKeyOption is a function that applies changes to a apiKeyConfig.
type APIKeyOption func(t *apiKeyConfig)

// WithAPIKeyName returns an option to set the API Key.
func WithAPIKeyName(apiKeyName, apiKeyPrivateKey string) APIKeyOption {
	return func(t *apiKeyConfig) {
		t.apiKeyName = apiKeyName
		t.apiKeyPrivateKey = apiKeyPrivateKey
	}
}

// WithLoadAPIKeyFromEnv returns an option to set whether to load the API
// Key from environment variables. If the API Key name and private key are both set,
// they take precedence over the environment variables.
func WithLoadAPIKeyFromEnv(loadAPIKeyFromEnv bool) APIKeyOption {
	return func(t *apiKeyConfig) {
		t.loadAPIKeyFromEnv = loadAPIKeyFromEnv
	}
}

// WithLoadAPIKeyFromFile returns an option to set whether to load the API
// Key from file directly. If the API Key name and private key are both set,
// they take precedence over the environment variables. Next if the env vars are
// set they take precedence or else the file is used if set.
func WithLoadAPIKeyFromFile(loadAPIKeyFromFile bool) APIKeyOption {
	return func(t *apiKeyConfig) {
		t.loadAPIKeyFromFile = loadAPIKeyFromFile
	}
}

// NewAPIKey creates a new Coinbase Cloud API Key based on the provided options.
func NewAPIKey(opts ...APIKeyOption) (*APIKey, error) {
	config := &apiKeyConfig{}

	for _, opt := range opts {
		opt(config)
	}

	if config.apiKeyName != "" && config.apiKeyPrivateKey != "" {
		return &APIKey{
			Name:       config.apiKeyName,
			PrivateKey: config.apiKeyPrivateKey,
		}, nil
	}

	if config.loadAPIKeyFromEnv {
		return loadAPIKeyFromEnv()
	}

	if config.loadAPIKeyFromFile {
		return loadAPIKeyFromFile()
	}

	return nil, fmt.Errorf("no valid api key gen option selected")
}

// loadAPIKeyFromEnv loads a new Coinbase Cloud API Key from environment variables.
func loadAPIKeyFromEnv() (*APIKey, error) {
	name := os.Getenv(nameEnvVar)
	if name == "" {
		return nil, errors.New("environment variable COINBASE_CLOUD_API_KEY_NAME must be set")
	}

	privateKey := os.Getenv(privateKeyEnvVar)
	if privateKey == "" {
		return nil, errors.New("environment variable COINBASE_CLOUD_API_KEY_PRIVATE_KEY must be set")
	}

	return &APIKey{
		Name:       name,
		PrivateKey: privateKey,
	}, nil
}

// loadAPIKeyFromFile loads a new Coinbase Cloud API Key from a private_key file downloaded from CB Cloud UI directly.
// It currently expects the private_key file to be located in path from where binary is being run.
func loadAPIKeyFromFile() (*APIKey, error) {
	keyFileName := ".coinbase_cloud_api_key.json"

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("file load: %w", err)
	}

	for wd != "" && wd != "/" {
		keyFilepath := path.Join(wd, keyFileName)
		wd = path.Dir(wd)

		f, err := os.Open(filepath.Clean(keyFilepath))
		if err != nil {
			// skip if file not accessible
			continue
		}

		var a APIKey

		dec := json.NewDecoder(f)
		if err := dec.Decode(&a); err != nil {
			return nil, fmt.Errorf("file load: %w", err)
		}

		return &a, nil
	}

	return nil, fmt.Errorf("error loading file")
}
