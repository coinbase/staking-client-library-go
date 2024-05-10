/*
 * This example code, demonstrates staking client library usage for listing supported Protocols, Networks and Actions.
 */

package main

import (
	"context"
	"errors"
	"log"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"google.golang.org/api/iterator"
)

var (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"
)

// An example function to demonstrate how to use the staking client libraries.
func main() {
	ctx := context.Background()

	// Loads the API key.
	apiKey, err := auth.NewAPIKey(auth.WithAPIKeyName(apiKeyName, apiPrivateKey))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	// Create a staking client.
	stakingClient, err := client.New(ctx, options.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	// List all workflows for a given project.
	workflowIter := stakingClient.Orchestration.ListWorkflows(ctx, &api.ListWorkflowsRequest{})

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
