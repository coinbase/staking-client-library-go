package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// PerformWorkflowStep helps update workflow move to the next state by returning the signed tx back.
func (s *StakingServiceClient) PerformWorkflowStep(
	ctx context.Context,
	req *stakingpb.PerformWorkflowStepRequest,
	opts ...gax.CallOption,
) (*stakingpb.Workflow, error) {
	wf, err := s.client.PerformWorkflowStep(ctx, req, opts...)
	if err == nil {
		return wf, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return wf, stakingerrors.FromError(err)
}
