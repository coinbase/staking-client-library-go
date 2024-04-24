package orchestration

import (
	"context"

	clients "github.com/coinbase/staking-client-library-go/client/options"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/orchestration/v1"
	"google.golang.org/grpc"
)

const (
	// ServiceName is the name of the service used by the Authenticator.
	serviceName = "staking"

	// ServiceEndpoint is the default endpoint URL to use.
	serviceEndpoint = "https://api.developer.coinbase.com/staking/orchestration"
)

// OrchestrationServiceClient is the client to use to access StakingService APIs.
type OrchestrationServiceClient struct {
	client *innerClient.StakingClient
}

// NewOrchestrationServiceClient returns a OrchestrationServiceClient based on the given inputs.
func NewOrchestrationServiceClient(
	ctx context.Context,
	stakingOpts ...clients.StakingClientOption,
) (*OrchestrationServiceClient, error) {
	config, err := clients.GetConfig(serviceName, serviceEndpoint, stakingOpts...)
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

	return &OrchestrationServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (s *OrchestrationServiceClient) Close() error {
	return s.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (s *OrchestrationServiceClient) Connection() *grpc.ClientConn {
	return s.client.Connection()
}
