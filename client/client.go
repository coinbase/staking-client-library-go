package client

import (
	"context"
	"fmt"

	"github.com/coinbase/staking-client-library-go/client/options"
	orchestrationClient "github.com/coinbase/staking-client-library-go/client/orchestration/v1"
	rewardsClient "github.com/coinbase/staking-client-library-go/client/rewards/v1"
)

type StakingClient struct {
	Orchestration *orchestrationClient.OrchestrationServiceClient
	Rewards       *rewardsClient.RewardsServiceClient
}

// New returns a StakingClient based on the given inputs.
func New(
	ctx context.Context,
	stakingOpts ...options.StakingClientOption,
) (*StakingClient, error) {
	orchestrationClient, err := orchestrationClient.NewOrchestrationServiceClient(ctx, stakingOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create orchestration client: %w", err)
	}

	rewardsClient, err := rewardsClient.NewRewardsServiceClient(ctx, stakingOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create rewards client: %w", err)
	}

	return &StakingClient{
		Orchestration: orchestrationClient,
		Rewards:       rewardsClient,
	}, nil
}
