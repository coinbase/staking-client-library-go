package orchestration

import (
	"context"
	"errors"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/orchestration/v1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// StakingTargetIterator is an interface for iterating through the response to ListStakingTargets.
type StakingTargetIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*stakingpb.StakingTarget, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *stakingpb.ListStakingTargetsResponse
}

// StakingTargetIteratorImpl is an implementation of StakingTargetIterator that unwraps correctly.
type StakingTargetIteratorImpl struct {
	iter *innerClient.StakingTargetIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *StakingTargetIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *StakingTargetIteratorImpl) Next() (*stakingpb.StakingTarget, error) {
	stakingTarget, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return stakingTarget, err
	}

	return stakingTarget, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *StakingTargetIteratorImpl) Response() *stakingpb.ListStakingTargetsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*stakingpb.ListStakingTargetsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListStakingTargets lists the StakingTargets supported by Staking API.
func (s *Client) ListStakingTargets(
	ctx context.Context,
	req *stakingpb.ListStakingTargetsRequest,
	opts ...gax.CallOption,
) StakingTargetIterator {
	return &StakingTargetIteratorImpl{iter: s.client.ListStakingTargets(ctx, req, opts...)}
}
