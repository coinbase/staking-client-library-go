package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// RefreshWorkflowStep helps refresh a workflow.
func (s *StakingServiceClient) RefreshWorkflowStep(
	ctx context.Context,
	req *stakingpb.RefreshWorkflowStepRequest,
	opts ...gax.CallOption,
) (*stakingpb.Workflow, error) {
	wf, err := s.client.RefreshWorkflowStep(ctx, req, opts...)
	if err == nil {
		return wf, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return wf, stakingerrors.FromError(err)
}
