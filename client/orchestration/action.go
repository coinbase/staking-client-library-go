package orchestration

import (
	"context"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
)

// ListActions lists the Actions supported by Staking API.
func (s *Client) ListActions(
	ctx context.Context,
	req *api.ListActionsRequest,
	opts ...gax.CallOption,
) (*api.ListActionsResponse, error) {
	actions, err := s.client.ListActions(ctx, req, opts...)
	if err != nil {
		err := stakingerrors.FromError(err)
		_ = err.Print()
	}

	return actions, err
}
