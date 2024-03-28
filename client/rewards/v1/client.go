package v1

import (
	"context"

	"google.golang.org/grpc"

	clients "github.com/coinbase/staking-client-library-go/client/options"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/rewards/v1"
)

const (
	// ServiceEndpoint is the endpoint the client will use.
	serviceEndpoint = "https://api.developer.coinbase.com/staking/rewards"
)

// RewardsServiceClient is the client to use to access StakingService APIs.
type RewardsServiceClient struct {
	client *innerClient.RewardClient
}

// NewRewardsServiceClient returns a RewardsServiceClient based on the given inputs.
func NewRewardsServiceClient(
	ctx context.Context,
	stakingOpts ...clients.StakingClientOption,
) (*RewardsServiceClient, error) {
	config, err := clients.GetConfig("rewards-reporting", serviceEndpoint, stakingOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewRewardRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	return &RewardsServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (s *RewardsServiceClient) Close() error {
	return s.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (s *RewardsServiceClient) Connection() *grpc.ClientConn {
	return s.client.Connection()
}
