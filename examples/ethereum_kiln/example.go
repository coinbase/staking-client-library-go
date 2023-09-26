/*
 * This example code, demonstrates staking client library usage for performing e2e staking on Ethereum Kiln.
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
	v1alpha1client "github.com/coinbase/staking-client-library-go/client/v1alpha1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
	"github.com/coinbase/staking-client-library-go/internal/signer"
)

const (
	// TODO: Replace with your project ID and private key.
	projectID  = ""
	privateKey = ""

	// TODO: Replace with your staker addresses and amount.
	stakerAddress             = ""
	integratorContractAddress = ""
	amount                    = "10"
	currency                  = "ETH"
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

	if projectID == "" || privateKey == "" || stakerAddress == "" {
		log.Fatalf("projectID, privateKey and stakerAddress must be set")
	}

	req := &stakingpb.CreateWorkflowRequest{
		Parent: fmt.Sprintf("projects/%s", projectID),
		Workflow: &stakingpb.Workflow{
			Action: "protocols/ethereum_kiln/networks/mainnet/actions/stake",
			StakingParameters: &stakingpb.Workflow_EthereumKilnStakingParameters{
				EthereumKilnStakingParameters: &stakingpb.EthereumKilnStakingParameters{
					Parameters: &stakingpb.EthereumKilnStakingParameters_StakeParameters{
						StakeParameters: &stakingpb.EthereumKilnStakeParameters{
							StakerAddress:             stakerAddress,
							IntegratorContractAddress: integratorContractAddress,
							Amount: &stakingpb.Amount{
								Value:    amount,
								Currency: currency,
							},
						},
					},
				},
			},
			SkipBroadcast: false,
		},
	}

	workflow, err := stakingClient.CreateWorkflow(ctx, req)
	if err != nil {
		sae := stakingerrors.FromError(err)
		_ = sae.Print()
		os.Exit(1)
	}

	log.Printf("Workflow created %s ...\n", workflow.Name)

	// Run loop until workflow reaches a terminal state
	for {
		// Get the latest workflow state
		workflow, err = stakingClient.GetWorkflow(ctx, &stakingpb.GetWorkflowRequest{Name: workflow.Name})
		if err != nil {
			log.Fatalf(fmt.Errorf("error getting workflow: %w", err).Error())
		}

		printWorkflowProgressDetails(workflow)

		// If workflow is in WAITING_FOR_SIGNING state, sign the transaction and update the workflow
		if v1alpha1client.WorkflowWaitingForSigning(workflow) {
			unsignedTx := workflow.Steps[workflow.GetCurrentStepId()].GetTxStepOutput().GetUnsignedTx()

			// Logic to sign the transaction. This can be substituted with any other signing mechanism.
			log.Printf("Signing unsigned tx %s ...\n", unsignedTx)

			signedTx, err := signer.New("ethereum_kiln").SignTransaction(privateKey, &signer.UnsignedTx{Data: []byte(unsignedTx)})
			if err != nil {
				log.Fatalf(fmt.Errorf("error signing transaction: %w", err).Error())
			}

			log.Printf("Returning back signed tx %s ...\n", string(signedTx.Data))

			workflow, err = stakingClient.PerformWorkflowStep(ctx, &stakingpb.PerformWorkflowStepRequest{
				Name: workflow.Name,
				Step: workflow.GetCurrentStepId(),
				Data: string(signedTx.Data),
			})
			if err != nil {
				log.Fatalf(fmt.Errorf("error updating workflow with signed tx: %w", err).Error())
			}
		} else if v1alpha1client.WorkflowWaitingForExternalBroadcast(workflow) {
			unsignedTx := workflow.Steps[workflow.GetCurrentStepId()].GetTxStepOutput().GetUnsignedTx()

			fmt.Printf("Please sign and broadcast this unsigned tx %s externally and return back the tx hash via the PerformWorkflowStep API ...\n", unsignedTx)
			break
		} else if v1alpha1client.WorkflowFailedRefreshable(workflow) {
			// If workflow is in FAILED_REFRESHABLE state, refresh the transaction and update the workflow
			fmt.Printf("Step %d failed but is refreshable, refreshing ...\n", workflow.GetCurrentStepId())

			workflow, err = stakingClient.RefreshWorkflowStep(ctx, &stakingpb.RefreshWorkflowStepRequest{
				Name: workflow.Name,
				Step: workflow.GetCurrentStepId(),
			})
			if err != nil {
				log.Fatalf(fmt.Errorf("error refreshing transaction: %w", err).Error())
			}
		} else if v1alpha1client.WorkflowFinished(workflow) {
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

	if v1alpha1client.WorkflowFinished(workflow) {
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
