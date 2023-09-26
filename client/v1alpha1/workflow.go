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

	return wf, sae
}

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

	return wf, sae
}

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

	return wf, sae
}

func WorkflowFinished(workflow *stakingpb.Workflow) bool {
	return workflow.State == stakingpb.Workflow_STATE_COMPLETED ||
		workflow.State == stakingpb.Workflow_STATE_FAILED ||
		workflow.State == stakingpb.Workflow_STATE_CANCELED ||
		workflow.State == stakingpb.Workflow_STATE_CANCEL_FAILED
}

func WorkflowWaitingForSigning(workflow *stakingpb.Workflow) bool {
	return workflow.State == stakingpb.Workflow_STATE_WAITING_FOR_SIGNING
}

func WorkflowWaitingForExternalBroadcast(workflow *stakingpb.Workflow) bool {
	return workflow.State == stakingpb.Workflow_STATE_WAITING_FOR_EXT_BROADCAST
}

func WorkflowFailedRefreshable(workflow *stakingpb.Workflow) bool {
	return workflow.State == stakingpb.Workflow_STATE_FAILED_REFRESHABLE
}
