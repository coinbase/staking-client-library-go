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
	"github.com/coinbase/staking-client-library-go/client/orchestration/v1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

const (
	// TODO: Replace with your project ID.
	projectID = ""

	// TODO: Replace with your wallet addresses and amount.
	walletAddress    = ""
	validatorAddress = "GkqYQysEGmuL6V2AJoNnWZUz2ZBGWhzQXsJiXm2CLKAN"
	amount           = "100000000"
	currency         = "SOL"
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

	if projectID == "" || walletAddress == "" {
		log.Fatalf("projectID and walletAddress must be set")
	}

	req := &stakingpb.CreateWorkflowRequest{
		Parent: fmt.Sprintf("projects/%s", projectID),
		Workflow: &stakingpb.Workflow{
			Action: "protocols/solana/networks/testnet/actions/stake",
			StakingParameters: &stakingpb.Workflow_SolanaStakingParameters{
				SolanaStakingParameters: &stakingpb.SolanaStakingParameters{
					Parameters: &stakingpb.SolanaStakingParameters_StakeParameters{
						StakeParameters: &stakingpb.SolanaStakeParameters{
							WalletAddress:    walletAddress,
							ValidatorAddress: validatorAddress,
							Amount: &stakingpb.Amount{
								Value:    amount,
								Currency: currency,
							},
						},
					},
				},
			},
			SkipBroadcast: true,
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
		workflow, err = stakingClient.Orchestration.GetWorkflow(ctx, &stakingpb.GetWorkflowRequest{Name: workflow.Name})
		if err != nil {
			log.Fatalf(fmt.Errorf("error getting workflow: %w", err).Error())
		}

		printWorkflowProgressDetails(workflow)

		// If workflow is in WAITING_FOR_SIGNING state, sign the transaction and update the workflow
		if v1.WorkflowWaitingForSigning(workflow) {
			unsignedTx := workflow.Steps[workflow.GetCurrentStepId()].GetTxStepOutput().GetUnsignedTx()

			fmt.Printf("Please sign this unsigned tx and return back the signed tx via the PerformWorkflowStep API : %s\n", unsignedTx)
			break
		} else if v1.WorkflowWaitingForExternalBroadcast(workflow) {
			unsignedTx := workflow.Steps[workflow.GetCurrentStepId()].GetTxStepOutput().GetUnsignedTx()

			fmt.Printf("Please sign and broadcast this unsigned tx externally and return back the tx hash via the PerformWorkflowStep API : %s\n", unsignedTx)
			break
		} else if v1.WorkflowFinished(workflow) {
			break
		}

		// Sleep for 5 seconds before polling for workflow status again
		time.Sleep(5 * time.Second)
	}
}

func printWorkflowProgressDetails(workflow *stakingpb.Workflow) {
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
	case *stakingpb.WorkflowStep_TxStepOutput:
		stepDetails = fmt.Sprintf("state: %s tx hash: %s",
			step.GetTxStepOutput().GetState().String(),
			step.GetTxStepOutput().GetTxHash(),
		)
	case *stakingpb.WorkflowStep_WaitStepOutput:
		stepDetails = fmt.Sprintf("state: %s current: %d target: %d",
			step.GetWaitStepOutput().GetState().String(),
			step.GetWaitStepOutput().GetCurrent(),
			step.GetWaitStepOutput().GetTarget(),
		)
	}

	if v1.WorkflowFinished(workflow) {
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
