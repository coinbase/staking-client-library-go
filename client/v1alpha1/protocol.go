package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// ListProtocols lists the Protocols supported by Staking API.
func (s *StakingServiceClient) ListProtocols(
	ctx context.Context,
	req *stakingpb.ListProtocolsRequest,
	opts ...gax.CallOption,
) (*stakingpb.ListProtocolsResponse, error) {
	protocols, err := s.client.ListProtocols(ctx, req, opts...)
	if err != nil {
		err := stakingerrors.FromError(err)
		_ = err.Print()
	}

	return protocols, err
}
