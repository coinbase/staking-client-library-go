package v1alpha1

import (
	"context"

	"github.com/googleapis/gax-go/v2"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

// CreateWorkflow starts a workflow with the given protocol specific parameters.
func (s *StakingServiceClient) CreateWorkflow(
	ctx context.Context,
	req *stakingpb.CreateWorkflowRequest,
	opts ...gax.CallOption,
) (*stakingpb.Workflow, error) {
	wf, err := s.client.CreateWorkflow(ctx, req, opts...)
	if err == nil {
		return wf, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return wf, sae
}
