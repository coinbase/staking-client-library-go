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

// NetworkIterator is an interface for iterating through the response to ListNetworks.
type NetworkIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*stakingpb.Network, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *stakingpb.ListNetworksResponse
}

// networkIteratorImpl is an implementation of NetworkIterator that unwraps correctly.
type networkIteratorImpl struct {
	iter *innerClient.NetworkIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *networkIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *networkIteratorImpl) Next() (*stakingpb.Network, error) {
	network, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return network, err
	}

	return network, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *networkIteratorImpl) Response() *stakingpb.ListNetworksResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*stakingpb.ListNetworksResponse)
	if !ok {
		return nil
	}

	return response
}

// ListNetworks lists the Networks supported by Staking API.
func (s *StakingServiceClient) ListNetworks(
	ctx context.Context,
	req *stakingpb.ListNetworksRequest,
	opts ...gax.CallOption,
) NetworkIterator {
	return &networkIteratorImpl{iter: s.client.ListNetworks(ctx, req, opts...)}
}
