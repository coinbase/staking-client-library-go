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

// An example function that verifies the API key has been configured correctly and you can connect to the Coinbase Staking API.
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

	// List all protocols.
	protocols, err := stakingClient.Orchestration.ListProtocols(ctx, &stakingpb.ListProtocolsRequest{})
	if err != nil {
		log.Fatalf("error listing protocols: %s", err.Error())
	}

	for _, protocol := range protocols.Protocols {
		log.Printf("got protocol: %s", protocol.Name)
	}

	protocol := "protocols/ethereum_kiln"

	// List all networks for a supported protocol.
	networks, err := stakingClient.Orchestration.ListNetworks(ctx, &stakingpb.ListNetworksRequest{
		Parent: protocol,
	})
	if err != nil {
		log.Fatalf("error listing networks: %s", err.Error())
	}

	for _, network := range networks.Networks {
		log.Printf("got network: %s", network.Name)
	}

	network := "protocols/ethereum_kiln/networks/holesky"

	// List all actions for a supported network.
	actions, err := stakingClient.Orchestration.ListActions(ctx, &stakingpb.ListActionsRequest{
		Parent: network,
	})
	if err != nil {
		log.Fatalf("error listing actions: %s", err.Error())
	}

	for _, action := range actions.Actions {
		log.Printf("got action: %s", action.Name)
	}

	// List all staking targets for a supported network.
	iter := stakingClient.Orchestration.ListStakingTargets(ctx, &stakingpb.ListStakingTargetsRequest{
		Parent: network,
	})

	for {
		stakingTarget, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing staking targets: %s", err.Error())
		}

		switch stakingTarget.GetStakingTargets().(type) {
		case *stakingpb.StakingTarget_Validator:
			log.Printf("got validator: %s", stakingTarget.GetValidator().String())
		case *stakingpb.StakingTarget_Contract:
			log.Printf("got contract: %s", stakingTarget.GetContract().String())
		}

	}
}
