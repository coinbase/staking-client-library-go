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

// ProtocolIterator is an interface for iterating through the response to ListProtocols.
type ProtocolIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*stakingpb.Protocol, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *stakingpb.ListProtocolsResponse
}

// protocolIteratorImpl is an implementation of ProtocolIterator that unwraps correctly.
type protocolIteratorImpl struct {
	iter *innerClient.ProtocolIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *protocolIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *protocolIteratorImpl) Next() (*stakingpb.Protocol, error) {
	protocol, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return protocol, err
	}

	return protocol, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *protocolIteratorImpl) Response() *stakingpb.ListProtocolsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*stakingpb.ListProtocolsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListProtocols lists the Protocols supported by Staking API.
func (s *StakingServiceClient) ListProtocols(
	ctx context.Context,
	req *stakingpb.ListProtocolsRequest,
	opts ...gax.CallOption,
) ProtocolIterator {
	return &protocolIteratorImpl{iter: s.client.ListProtocols(ctx, req, opts...)}
}
