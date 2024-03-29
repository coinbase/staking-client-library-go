package v1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

// ListNetworks lists the Networks supported by Staking API.
func (s *OrchestrationServiceClient) ListNetworks(
	ctx context.Context,
	req *stakingpb.ListNetworksRequest,
	opts ...gax.CallOption,
) (*stakingpb.ListNetworksResponse, error) {
	networks, err := s.client.ListNetworks(ctx, req, opts...)
	if err != nil {
		err := stakingerrors.FromError(err)
		_ = err.Print()
	}

	return networks, err
}
