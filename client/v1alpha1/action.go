package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// ListActions lists the Actions supported by Staking API.
func (s *StakingServiceClient) ListActions(
	ctx context.Context,
	req *stakingpb.ListActionsRequest,
	opts ...gax.CallOption,
) (*stakingpb.ListActionsResponse, error) {
	actions, err := s.client.ListActions(ctx, req, opts...)
	if err != nil {
		err := stakingerrors.FromError(err)
		_ = err.Print()
	}

	return actions, err
}
