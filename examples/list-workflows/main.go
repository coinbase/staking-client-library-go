/*
 * This example code, demonstrates staking client library usage for listing supported Protocols, Networks and Actions.
 */

package main

import (
	"context"
	"errors"
	"log"

	"google.golang.org/api/iterator"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

// An example function to demonstrate how to use the staking client libraries.
func main() {
	ctx := context.Background()

	apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	authOpt := options.WithAPIKey(apiKey)

	// Create a staking client.
	stakingClient, err := client.New(ctx, authOpt)
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	// List all workflows for a given project.
	workflowIter := stakingClient.Orchestration.ListWorkflows(ctx, &stakingpb.ListWorkflowsRequest{})

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
