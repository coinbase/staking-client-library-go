package v1alpha1

import (
	"context"
	"errors"

	"github.com/googleapis/gax-go/v2"

	"google.golang.org/api/iterator"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/v1alpha1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// ActionIterator is an interface for iterating through the response to ListActions.
type ActionIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*stakingpb.Action, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *stakingpb.ListActionsResponse
}

// ActionIteratorImpl is an implementation of ActionIterator that unwraps correctly.
type ActionIteratorImpl struct {
	iter *innerClient.ActionIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *ActionIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *ActionIteratorImpl) Next() (*stakingpb.Action, error) {
	action, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return action, err
	}

	return action, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *ActionIteratorImpl) Response() *stakingpb.ListActionsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*stakingpb.ListActionsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListActions lists the Actions supported by Staking API.
func (s *StakingServiceClient) ListActions(
	ctx context.Context,
	req *stakingpb.ListActionsRequest,
	opts ...gax.CallOption,
) ActionIterator {
	return &ActionIteratorImpl{iter: s.client.ListActions(ctx, req, opts...)}
}
