/*
 * This example code, demonstrates staking client library usage for performing e2e staking on Polygon.
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/iterator"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	v1alpha1client "github.com/coinbase/staking-client-library-go/client/v1alpha1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
	"github.com/coinbase/staking-client-library-go/internal/signer"
)

const (
	// TODO: Replace with your project ID and private key.
	projectID  = ""
	privateKey = ""

	// TODO: Replace with your delegator addresses and amount.
	delegatorAddress = ""
	validatorAddress = "0x857679d69fe50e7b722f94acd2629d80c355163d"
	amount           = "1"
	currency         = "MATIC"
)

// An example function to demonstrate how to use the staking client libraries.
func main() {
	ctx := context.Background()

	apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
	if err != nil {
		log.Fatalf("error loading API key: %v", err)
	}

	authOpt := client.WithAPIKey(apiKey)

	// Create a staking client.
	stakingClient, err := v1alpha1client.NewStakingServiceClient(ctx, authOpt)
	if err != nil {
		log.Fatalf("error instantiating staking client: %v", err)
	}

	// List all protocols.
	protocolIter := stakingClient.ListProtocols(ctx, &stakingpb.ListProtocolsRequest{})

	for {
		protocol, err := protocolIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing protocols: %v", err)
		}

		log.Printf("got protocol: %v", protocol.Name)
	}

	protocol := "protocols/polygon"

	// List all networks for a supported protocol.
	networkIter := stakingClient.ListNetworks(ctx, &stakingpb.ListNetworksRequest{
		Parent: protocol,
	})

	for {
		network, err := networkIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing networks: %v", err.Error())
		}

		log.Printf("got network: %v", network.Name)
	}

	network := "protocols/polygon/networks/mainnet"

	// List all actions for a supported network.
	actionIter := stakingClient.ListActions(ctx, &stakingpb.ListActionsRequest{
		Parent: network,
	})

	for {
		action, err := actionIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing actions: %v", err)
		}

		log.Printf("got action: %v", action.Name)
	}

	// List all validators for a supported network.
	validatorIter := stakingClient.ListValidators(ctx, &stakingpb.ListValidatorsRequest{
		Parent: network,
	})

	for {
		validator, err := validatorIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing validators: %v", err)
		}

		log.Printf("got validator: %v", validator.Name)
	}

	if projectID == "" || privateKey == "" || delegatorAddress == "" {
		log.Fatalf("projectID, privateKey and delegatorAddress must be set")
	}

	req := &stakingpb.CreateWorkflowRequest{
		Parent: fmt.Sprintf("projects/%s", projectID),
		Workflow: &stakingpb.Workflow{
			Action: "protocols/polygon/networks/mainnet/actions/stake",
			StakingParameters: &stakingpb.Workflow_PolygonStakingParameters{
				PolygonStakingParameters: &stakingpb.PolygonStakingParameters{
					Parameters: &stakingpb.PolygonStakingParameters_StakeParameters{
						StakeParameters: &stakingpb.PolygonStakeParameters{
							DelegatorAddress: delegatorAddress,
							ValidatorAddress: validatorAddress,
							Amount: &stakingpb.Amount{
								Value:    amount,
								Currency: currency,
							},
						},
					},
				},
			},
		},
	}

	// Create a workflow for Polygon staking
	var workflow *stakingpb.Workflow

	workflow, err = stakingClient.CreateWorkflow(ctx, req)
	if err != nil {
		log.Fatalf("errror creating workflow: %v", err.Error())
	}

	log.Printf("Workflow created %s ...\n", workflow.Name)

	workflow, err = stakingClient.GetWorkflow(ctx, &stakingpb.GetWorkflowRequest{Name: workflow.Name})
	if err != nil {
		log.Fatalf("error getting workflow: %v", err)
	}

	// Run loop until workflow reaches terminal state of COMPLETED or FAILED
	for workflow.State != stakingpb.State_COMPLETED && workflow.State != stakingpb.State_FAILED {
		step := workflow.Steps[workflow.GetCurrentStepId()]
		// If workflow is in WAITING_FOR_SIGNING state, sign the transaction and update the workflow
		if workflow.State == stakingpb.State_WAITING_FOR_SIGNING {
			unsignedTx := step.GetTxStepOutput().GetUnsignedTx()

			// Logic to sign the transaction. This can be substituted with any other signing mechanism.
			log.Printf("Signing unsigned tx %s ...\n", unsignedTx)

			signedTx, err := signer.New("polygon").SignTransaction(privateKey, &signer.UnsignedTx{Data: []byte(unsignedTx)})
			if err != nil {
				log.Fatalf("error signing transaction: %v", err)
			}

			log.Printf("Returning back signed tx %s ...\n", string(signedTx.Data))

			workflow, err = stakingClient.PerformWorkflowStep(ctx, &stakingpb.PerformWorkflowStepRequest{
				Name:     workflow.Name,
				Step:     workflow.GetCurrentStepId(),
				SignedTx: string(signedTx.Data),
			})
			if err != nil {
				log.Fatalf("error updating workflow with signed tx: %v", err)
			}
		} else if workflow.State == stakingpb.State_IN_PROGRESS {
			// If workflow is in IN_PROGRESS state, print the details of the current step
			switch output := step.Output.(type) {
			case *stakingpb.WorkflowStep_TxStepOutput:
				txHash := output.TxStepOutput.GetTxHash()
				if txHash != "" {
					log.Println("Tx Hash: ", txHash)
				}
			case *stakingpb.WorkflowStep_WaitStepOutput:
				target := output.WaitStepOutput.GetTarget()
				current := output.WaitStepOutput.GetCurrent()
				unit := output.WaitStepOutput.GetUnit()
				log.Println("Waiting for ", (target-current)/10_000, " ", unit, " to complete ...")
			default:
				panic("unknown step type encountered")
			}
		}

		// Get the latest workflow state
		workflow, err = stakingClient.GetWorkflow(ctx, &stakingpb.GetWorkflowRequest{Name: workflow.Name})
		if err != nil {
			log.Fatalf("error getting workflow: %v", err)
		}

		createTime := workflow.GetCreateTime().AsTime()
		updateTime := workflow.GetUpdateTime().AsTime()
		log.Printf("Waiting for workflow to finish - name: %s state: %s runtime: %v\n", step.GetName(), step.GetState(), updateTime.Sub(createTime))

		// Sleep for 5 seconds before polling for workflow status again
		time.Sleep(5 * time.Second)
	}
}
