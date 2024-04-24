package orchestration

import (
	"context"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
)

// ListNetworks lists the Networks supported by Staking API.
func (s *Client) ListNetworks(
	ctx context.Context,
	req *api.ListNetworksRequest,
	opts ...gax.CallOption,
) (*api.ListNetworksResponse, error) {
	networks, err := s.client.ListNetworks(ctx, req, opts...)
	if err != nil {
		err := stakingerrors.FromError(err)
		_ = err.Print()
	}

	return networks, err
}
