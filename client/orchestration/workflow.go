package orchestration

import (
	"context"
	"errors"

	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	innerClient "github.com/coinbase/staking-client-library-go/gen/client/coinbase/staking/orchestration/v1"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// CreateWorkflow starts a workflow with the given protocol specific parameters.
func (s *Client) CreateWorkflow(
	ctx context.Context,
	req *api.CreateWorkflowRequest,
	opts ...gax.CallOption,
) (*api.Workflow, error) {
	wf, err := s.client.CreateWorkflow(ctx, req, opts...)
	if err == nil {
		return wf, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return wf, sae
}

// GetWorkflow get the current state of a workflow.
func (s *Client) GetWorkflow(
	ctx context.Context,
	req *api.GetWorkflowRequest,
	opts ...gax.CallOption,
) (*api.Workflow, error) {
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
	Next() (*api.Workflow, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *api.ListWorkflowsResponse
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
func (n *WorkflowIteratorImpl) Next() (*api.Workflow, error) {
	workflow, err := n.iter.Next()
	if errors.Is(err, iterator.Done) || err == nil {
		return workflow, err
	}

	return workflow, stakingerrors.FromError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *WorkflowIteratorImpl) Response() *api.ListWorkflowsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response, ok := n.iter.Response.(*api.ListWorkflowsResponse)
	if !ok {
		return nil
	}

	return response
}

// ListWorkflows lists the Workflows supported by Staking API.
func (s *Client) ListWorkflows(
	ctx context.Context,
	req *api.ListWorkflowsRequest,
	opts ...gax.CallOption,
) WorkflowIterator {
	return &WorkflowIteratorImpl{iter: s.client.ListWorkflows(ctx, req, opts...)}
}

// PerformWorkflowStep helps update workflow move to the next state by returning the signed tx back.
func (s *Client) PerformWorkflowStep(
	ctx context.Context,
	req *api.PerformWorkflowStepRequest,
	opts ...gax.CallOption,
) (*api.Workflow, error) {
	wf, err := s.client.PerformWorkflowStep(ctx, req, opts...)
	if err == nil {
		return wf, nil
	}

	sae := stakingerrors.FromError(err)
	_ = sae.Print()

	return wf, sae
}

func WorkflowFinished(workflow *api.Workflow) bool {
	return workflow.State == api.Workflow_STATE_COMPLETED ||
		workflow.State == api.Workflow_STATE_FAILED
}

func WorkflowWaitingForExternalBroadcast(workflow *api.Workflow) bool {
	return workflow.State == api.Workflow_STATE_WAITING_FOR_EXT_BROADCAST
}
