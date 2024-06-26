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

// StakeIterator is an interface for iterating through the response to ListStakes.
type StakeIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*api.Stake, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *api.ListStakesResponse
}

// StakeIteratorImpl is an implementation of StakeIterator that unwraps correctly.
type StakeIteratorImpl struct {
	iter *innerClient.StakeIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *StakeIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *StakeIteratorImpl) Next() (*api.Stake, error) {
	reward, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return reward, err
	}

	return reward, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *StakeIteratorImpl) Response() *api.ListStakesResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*api.ListStakesResponse)
	if !ok {
		return nil
	}

	return response
}

// ListStakes list staking activities for a given protocol.
func (s *Client) ListStakes(
	ctx context.Context,
	req *api.ListStakesRequest,
	opts ...gax.CallOption,
) StakeIterator {
	return &StakeIteratorImpl{iter: s.client.ListStakes(ctx, req, opts...)}
}
