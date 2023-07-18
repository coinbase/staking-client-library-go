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

// ValidatorIterator is an interface for iterating through the response to ListValidators.
type ValidatorIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*stakingpb.Validator, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *stakingpb.ListValidatorsResponse
}

// ValidatorIteratorImpl is an implementation of ValidatorIterator that unwraps correctly.
type ValidatorIteratorImpl struct {
	iter *innerClient.ValidatorIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *ValidatorIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *ValidatorIteratorImpl) Next() (*stakingpb.Validator, error) {
	validator, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return validator, err
	}

	return validator, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *ValidatorIteratorImpl) Response() *stakingpb.ListValidatorsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*stakingpb.ListValidatorsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListValidators lists the Validators supported by Staking API.
func (s *StakingServiceClient) ListValidators(
	ctx context.Context,
	req *stakingpb.ListValidatorsRequest,
	opts ...gax.CallOption,
) ValidatorIterator {
	return &ValidatorIteratorImpl{iter: s.client.ListValidators(ctx, req, opts...)}
}
