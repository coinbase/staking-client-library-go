package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// GetWorkflow get the current state of a workflow.
func (s *StakingServiceClient) GetWorkflow(
	ctx context.Context,
	req *stakingpb.GetWorkflowRequest,
	opts ...gax.CallOption,
) (*stakingpb.Workflow, error) {
	wf, err := s.client.GetWorkflow(ctx, req, opts...)
	if err == nil {
		return wf, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return wf, stakingerrors.FromError(err)
}
