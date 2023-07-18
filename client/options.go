package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/googleapis/gax-go/v2"

	"google.golang.org/api/option"

	"github.com/coinbase/staking-client-library-go/auth"
)

// StakingClientConfig stores configuration information about a Staking client.
type StakingClientConfig struct {
	Endpoint      string
	ServiceName   string
	APIKey        *auth.APIKey
	HTTPClient    *http.Client
	ClientOptions []option.ClientOption
}

// StakingClientOption is a function that applies changes to a clientConfig.
type StakingClientOption func(c *StakingClientConfig)

// WithEndpoint returns an option to set the service endpoint.
func WithEndpoint(endpoint string) StakingClientOption {
	return func(c *StakingClientConfig) {
		c.Endpoint = endpoint
	}
}

// WithAPIKey returns an option to set the Coinbase Cloud API Key.
func WithAPIKey(apiKey *auth.APIKey) StakingClientOption {
	return func(c *StakingClientConfig) {
		c.APIKey = apiKey
	}
}

// WithHTTPClient returns an option to set the HTTP Client.
func WithHTTPClient(httpClient *http.Client) StakingClientOption {
	return func(c *StakingClientConfig) {
		c.HTTPClient = httpClient
	}
}

// WithClientOptions returns an option to set the inner set of ClientOptions.
func WithClientOptions(clientOptions ...option.ClientOption) StakingClientOption {
	return func(c *StakingClientConfig) {
		c.ClientOptions = clientOptions
	}
}

// GetConfig returns a StakingClientConfig with all the StakingClientOptions applied.
// It uses the given default endpoint if none of the options set it.
func GetConfig(
	serviceName string,
	defaultEndpoint string,
	stakingOpts ...StakingClientOption,
) (*StakingClientConfig, error) {
	config := &StakingClientConfig{}

	for _, stakingOpt := range stakingOpts {
		stakingOpt(config)
	}

	if config.Endpoint == "" {
		config.Endpoint = defaultEndpoint
	}

	if serviceName == "" {
		return nil, errors.New("service name must be provided")
	}

	config.ServiceName = serviceName

	return config, nil
}

// GetClientOptions returns the list of ClientOptions based on the given endpoint, service name,
// and config.
func GetClientOptions(config *StakingClientConfig) ([]option.ClientOption, error) {
	clientOptions := []option.ClientOption{}

	// Start with explicitly indicated client options.
	if len(config.ClientOptions) > 0 {
		clientOptions = append(clientOptions, config.ClientOptions...)
	}

	if config.Endpoint == "" {
		return nil, errors.New("endpoint must be provided")
	}

	clientOptions = append(clientOptions, option.WithEndpoint(config.Endpoint))

	if config.ServiceName == "" {
		return nil, errors.New("service name must be provided")
	}

	httpClient, err := GetHTTPClient(config.ServiceName, config)
	if err != nil {
		return nil, err
	}

	clientOptions = append(clientOptions, option.WithHTTPClient(httpClient))

	return clientOptions, nil
}

// GetHTTPClient returns the HTTPClient based on the given service name and config.
func GetHTTPClient(serviceName string, config *StakingClientConfig) (*http.Client, error) {
	var httpClient *http.Client
	if config.HTTPClient != nil {
		httpClient = config.HTTPClient
	} else {
		httpClient = &http.Client{}
	}

	if config.APIKey != nil {
		if httpClient.Transport == nil {
			httpClient.Transport = http.DefaultTransport
		}

		transport, err := NewTransport(
			httpClient.Transport,
			serviceName,
			config.APIKey,
		)
		if err != nil {
			return nil, err
		}

		httpClient.Transport = transport
	}

	return httpClient, nil
}

// LongRunningOperation is the interface for long-running operations that is
// used to create the gax call options for interacting with LROs.
type LongRunningOperation interface {
	Name() string
	PathPrefix() string
}

// LROOptions returns the call options for long-running operations.
// This overrides the gapic generated client `WithPath` call option that ignores the
// path prefix, with a call option that includes the path prefix.
func LROOptions(op LongRunningOperation, version string, opts []gax.CallOption) []gax.CallOption {
	if op.PathPrefix() == "" {
		return opts
	}

	return append(opts, gax.WithPath(fmt.Sprintf("%s/%s/%s", op.PathPrefix(), version, op.Name())))
}
