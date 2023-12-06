/*
 * This example code, demonstrates staking client library usage for listing supported Protocols, Networks and Actions.
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	v1alpha1client "github.com/coinbase/staking-client-library-go/client/v1alpha1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
	"google.golang.org/api/iterator"
)

const (
	// TODO: Replace with your project ID.
	projectID = ""
)

// An example function to demonstrate how to use the staking client libraries.
func main() {
	ctx := context.Background()

	apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	authOpt := client.WithAPIKey(apiKey)

	// Create a staking client.
	stakingClient, err := v1alpha1client.NewStakingServiceClient(ctx, authOpt)
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	if projectID == "" {
		log.Fatalf("projectID must be set")
	}

	// List all workflows for a given project.
	workflowIter := stakingClient.ListWorkflows(ctx, &stakingpb.ListWorkflowsRequest{
		Parent: fmt.Sprintf("projects/%s", projectID),
	})

	for {
		workflow, err := workflowIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing workflows: %s", err.Error())
		}

		log.Println("workflow name:", workflow.GetName())
	}
}
