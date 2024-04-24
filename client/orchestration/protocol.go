package orchestration

import (
	"context"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
)

// ListProtocols lists the Protocols supported by Staking API.
func (s *Client) ListProtocols(
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
