package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// ViewStakingContext helps view staking context information given a specific network address.
func (s *StakingServiceClient) ViewStakingContext(
	ctx context.Context,
	req *stakingpb.ViewStakingContextRequest,
	opts ...gax.CallOption,
) (*stakingpb.ViewStakingContextResponse, error) {
	response, err := s.client.ViewStakingContext(ctx, req, opts...)
	if err == nil {
		return response, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return nil, sae
}
