/*
 * This example code, demonstrates staking client library usage for performing e2e staking on Polygon.
 */

package main

import (
	"context"
	"errors"
	"google.golang.org/api/iterator"
	"log"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	v1alpha1client "github.com/coinbase/staking-client-library-go/client/v1alpha1"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/v1alpha1"
)

const (
	// TODO: Replace with your project ID and private key.
	projectID = ""
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
}
