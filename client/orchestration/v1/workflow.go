package v1

import (
	"context"
	"errors"

	"github.com/googleapis/gax-go/v2"

	"google.golang.org/api/iterator"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/orchestration/v1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

// CreateWorkflow starts a workflow with the given protocol specific parameters.
func (s *OrchestrationServiceClient) CreateWorkflow(
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
func (s *OrchestrationServiceClient) GetWorkflow(
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

// WorkflowIterator is an interface for iterating through the response to ListWorkflows.
type WorkflowIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*stakingpb.Workflow, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *stakingpb.ListWorkflowsResponse
}

// WorkflowIteratorImpl is an implementation of WorkflowIterator that unwraps correctly.
type WorkflowIteratorImpl struct {
	iter *innerClient.WorkflowIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *WorkflowIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *WorkflowIteratorImpl) Next() (*stakingpb.Workflow, error) {
	workflow, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return workflow, err
	}

	return workflow, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *WorkflowIteratorImpl) Response() *stakingpb.ListWorkflowsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*stakingpb.ListWorkflowsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListWorkflows lists the Workflows supported by Staking API.
func (s *OrchestrationServiceClient) ListWorkflows(
	ctx context.Context,
	req *stakingpb.ListWorkflowsRequest,
	opts ...gax.CallOption,
) WorkflowIterator {
	return &WorkflowIteratorImpl{iter: s.client.ListWorkflows(ctx, req, opts...)}
}

// PerformWorkflowStep helps update workflow move to the next state by returning the signed tx back.
func (s *OrchestrationServiceClient) PerformWorkflowStep(
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

func WorkflowFinished(workflow *stakingpb.Workflow) bool {
	return workflow.State == stakingpb.Workflow_STATE_COMPLETED ||
		workflow.State == stakingpb.Workflow_STATE_FAILED
}

func WorkflowWaitingForExternalBroadcast(workflow *stakingpb.Workflow) bool {
	return workflow.State == stakingpb.Workflow_STATE_WAITING_FOR_EXT_BROADCAST
}
