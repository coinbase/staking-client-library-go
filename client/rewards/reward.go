package rewards

import (
	"context"
	"errors"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/rewards/v1"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// RewardIterator is an interface for iterating through the response to ListRewards.
type RewardIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*api.Reward, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *api.ListRewardsResponse
}

// RewardIteratorImpl is an implementation of RewardIterator that unwraps correctly.
type RewardIteratorImpl struct {
	iter *innerClient.RewardIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *RewardIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *RewardIteratorImpl) Next() (*api.Reward, error) {
	reward, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return reward, err
	}

	return reward, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *RewardIteratorImpl) Response() *api.ListRewardsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*api.ListRewardsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListRewards helps list onchain rewards of an address for a specific protocol, with optional filters for time range, aggregation period, and more.
func (s *Client) ListRewards(
	ctx context.Context,
	req *api.ListRewardsRequest,
	opts ...gax.CallOption,
) RewardIterator {
	return &RewardIteratorImpl{iter: s.client.ListRewards(ctx, req, opts...)}
}
