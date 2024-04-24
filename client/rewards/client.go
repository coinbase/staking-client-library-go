package rewards

import (
	"context"

	"github.com/coinbase/staking-client-library-go/client/options"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/rewards/v1"
	"google.golang.org/grpc"
)

const (
	// ServiceEndpoint is the endpoint the client will use.
	serviceEndpoint = "https://api.developer.coinbase.com/staking/rewards"
)

// Client is the client to use to access StakingService APIs.
type Client struct {
	client *innerClient.RewardClient
}

// NewRewardsServiceClient returns a RewardsServiceClient based on the given inputs.
func NewRewardsServiceClient(
	ctx context.Context,
	stakingOpts ...options.StakingClientOption,
) (*Client, error) {
	config, err := options.GetConfig("rewards-reporting", serviceEndpoint, stakingOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := options.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewRewardRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (s *Client) Close() error {
	return s.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (s *Client) Connection() *grpc.ClientConn {
	return s.client.Connection()
}
