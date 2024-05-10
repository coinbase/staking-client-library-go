/*
 * This example code, demonstrates staking client library usage for performing e2e staking on Solana.
 */

package main

import (
	"context"
	"log"
	"os"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	"github.com/coinbase/staking-client-library-go/client/options"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	// TODO: Replace with your wallet addresses and amount.
	walletAddress = "8rMGARtkJY5QygP1mgvBFLsE9JrvXByARJiyNfcSE5Z"
	amount        = "100000000"
	currency      = "SOL"
)

// An example function to demonstrate how to use the staking client libraries.
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
			Action: "protocols/solana/networks/devnet/actions/stake",
			StakingParameters: &api.Workflow_SolanaStakingParameters{
				SolanaStakingParameters: &api.SolanaStakingParameters{
					Parameters: &api.SolanaStakingParameters_StakeParameters{
						StakeParameters: &api.SolanaStakeParameters{
							WalletAddress: walletAddress,
							Amount: &api.Amount{
								Value:    amount,
								Currency: currency,
							},
						},
					},
				},
			},
		},
	}

	workflow, err := stakingClient.Orchestration.CreateWorkflow(ctx, req)
	if err != nil {
		sae := stakingerrors.FromError(err)
		_ = sae.Print()
		os.Exit(1)
	}

	marshaled, err := protojson.MarshalOptions{Indent: "  ", Multiline: true}.Marshal(workflow)
	if err != nil {
		log.Fatalf("error marshaling reward: %s", err.Error())
	}

	log.Printf("Workflow created: \n%s", marshaled)
}
