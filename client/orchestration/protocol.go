package orchestration

import (
	"context"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
)

// ListProtocols lists the Protocols supported by Staking API.
func (s *Client) ListProtocols(
	ctx context.Context,
	req *api.ListProtocolsRequest,
	opts ...gax.CallOption,
) (*api.ListProtocolsResponse, error) {
	protocols, err := s.client.ListProtocols(ctx, req, opts...)
	if err != nil {
		err := stakingerrors.FromError(err)
		_ = err.Print()
	}

	return protocols, err
}
