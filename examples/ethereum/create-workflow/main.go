/*
 * This example code demonstrates creating an Ethereum Partial ETH (<32 ETH) workflow, the fundamental process handler for an E2E staking experience.
 */

package main

import (
	"context"
	"log"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

func main() {
	ctx := context.Background()

	// Loads the API key from the default location.
	apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	// Creates the Coinbase Staking API client.
	stakingClient, err := client.New(ctx, options.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	req := &api.CreateWorkflowRequest{
		Workflow: &api.Workflow{
			Action: "protocols/ethereum_kiln/networks/holesky/actions/stake",
			StakingParameters: &api.Workflow_EthereumKilnStakingParameters{
				EthereumKilnStakingParameters: &api.EthereumKilnStakingParameters{
					Parameters: &api.EthereumKilnStakingParameters_StakeParameters{
						StakeParameters: &api.EthereumKilnStakeParameters{
							StakerAddress: "0xdb816889F2a7362EF242E5a717dfD5B38Ae849FE",
							Amount: &api.Amount{
								Value:    "20",
								Currency: "ETH",
							},
						},
					},
				},
			},
		},
	}

	workflow, err := stakingClient.Orchestration.CreateWorkflow(ctx, req)
	if err != nil {
		log.Fatalf("couldn't create workflow: %s", err.Error())
	}

	marshaled, err := protojson.MarshalOptions{Indent: "  ", Multiline: true}.Marshal(workflow)
	if err != nil {
		log.Fatalf("error marshaling reward: %s", err.Error())
	}

	log.Printf("Workflow created: \n%s", marshaled)
}
