package v1alpha1

import (
	"context"

	"google.golang.org/grpc"

	clients "github.com/coinbase/staking-client-library-go/client"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/v1alpha1"
)

const (
	// StakingServiceName is the name of the ProtocolService used by the Authenticator.
	stakingServiceName = "staking"

	// StakingServiceEndpoint is the default endpoint URL to use for ProtocolService.
	stakingServiceEndpoint = "https://api.developer.coinbase.com/staking"
)

// StakingServiceClient is the client to use to access StakingService APIs.
type StakingServiceClient struct {
	client *innerClient.StakingClient
}

// NewStakingServiceClient returns a StakingServiceClient based on the given inputs.
func NewStakingServiceClient(
	ctx context.Context,
	stakingOpts ...clients.StakingClientOption,
) (*StakingServiceClient, error) {
	config, err := clients.GetConfig(stakingServiceName, stakingServiceEndpoint, stakingOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewStakingRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	return &StakingServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (s *StakingServiceClient) Close() error {
	return s.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (s *StakingServiceClient) Connection() *grpc.ClientConn {
	return s.client.Connection()
}
