package orchestration

import (
	"context"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
)

// ViewStakingContext helps view staking context information given a specific network address.
func (s *Client) ViewStakingContext(
	ctx context.Context,
	req *api.ViewStakingContextRequest,
	opts ...gax.CallOption,
) (*api.ViewStakingContextResponse, error) {
	response, err := s.client.ViewStakingContext(ctx, req, opts...)
	if err == nil {
		return response, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return nil, sae
}
