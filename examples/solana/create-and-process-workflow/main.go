/*
 * This example code, demonstrates staking client library usage for performing e2e staking on Solana.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	stakingerrors "github.com/coinbase/staking-client-library-go/client/errors"
	"github.com/coinbase/staking-client-library-go/client/options"
	"github.com/coinbase/staking-client-library-go/client/orchestration"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
	"github.com/coinbase/staking-client-library-go/internal/signer"
)

const (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"

	// TODO: Replace with your wallet's private key.
	walletPrivateKey = ""

	// TODO: Replace with your wallet addresses and amount.
	walletAddress = ""
	amount        = "100000000"
	currency      = "SOL"
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

	if walletPrivateKey == "" || walletAddress == "" {
		log.Fatalf("walletPrivateKey and walletAddress must be set")
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

	log.Printf("Workflow created %s ...\n", workflow.Name)

	// Run loop until workflow reaches a terminal state
	for {
		// Get the latest workflow state
		workflow, err = stakingClient.Orchestration.GetWorkflow(ctx, &api.GetWorkflowRequest{Name: workflow.Name})
		if err != nil {
			log.Fatalf(fmt.Errorf("error getting workflow: %w", err).Error())
		}

		printWorkflowProgressDetails(workflow)

		// If workflow is in WAITING_FOR_EXT_BROADCAST state, sign, broadcast the transaction and update the workflow.
		if orchestration.WorkflowWaitingForExternalBroadcast(workflow) {
			unsignedTx := workflow.Steps[workflow.GetCurrentStepId()].GetTxStepOutput().GetUnsignedTx()

			// Logic to sign the transaction. This can be substituted with any other signing mechanism.
			log.Printf("Signing unsigned tx %s ...\n", unsignedTx)

			signedTx, err := signer.New("solana").SignTransaction([]string{walletPrivateKey}, &signer.UnsignedTx{Data: []byte(unsignedTx)})
			if err != nil {
				log.Fatalf(fmt.Errorf("error signing transaction: %w", err).Error())
			}

			// Add logic to broadcast the tx here.
			fmt.Printf("Please broadcast this signed tx %s externally and return back the tx hash via the PerformWorkflowStep API ...\n", signedTx)
			break
		} else if orchestration.WorkflowFinished(workflow) {
			break
		}

		// Sleep for 1 second before polling for workflow status again
		time.Sleep(1 * time.Second)
	}
}

func printWorkflowProgressDetails(workflow *api.Workflow) {
	if len(workflow.GetSteps()) <= 0 {
		fmt.Println("Waiting for steps to be created ...")
		time.Sleep(2 * time.Second)
	}

	step := workflow.Steps[workflow.GetCurrentStepId()]

	createTime := workflow.GetCreateTime().AsTime()
	updateTime := workflow.GetUpdateTime().AsTime()
	runtime := updateTime.Sub(createTime)

	var stepDetails string

	switch step.GetOutput().(type) {
	case *api.WorkflowStep_TxStepOutput:
		stepDetails = fmt.Sprintf("state: %s tx hash: %s",
			step.GetTxStepOutput().GetState().String(),
			step.GetTxStepOutput().GetTxHash(),
		)
	case *api.WorkflowStep_WaitStepOutput:
		stepDetails = fmt.Sprintf("state: %s current: %d target: %d",
			step.GetWaitStepOutput().GetState().String(),
			step.GetWaitStepOutput().GetCurrent(),
			step.GetWaitStepOutput().GetTarget(),
		)
	}

	if orchestration.WorkflowFinished(workflow) {
		log.Printf("Workflow reached end state - step name: %s %s workflow state: %s runtime: %v\n",
			step.GetName(),
			stepDetails,
			workflow.GetState().String(),
			runtime,
		)
	} else {
		log.Printf("Waiting for workflow to finish - step name: %s %s workflow state: %s runtime: %v\n",
			step.GetName(),
			stepDetails,
			workflow.GetState().String(),
			runtime,
		)
	}
}
