package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	"github.com/coinbase/staking-client-library-go/client/rewards"
	filter "github.com/coinbase/staking-client-library-go/client/rewards/stakesfilter"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

/*
 * Run the code with 'go run examples/cosmos/list-stakes/main.go' to view the balances for Coinbase Cloud's public validator.
 * Or, to view balances for any arbitrary validator, simply replace the address below with any validator on the Cosmos blockchain.
 */

const (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"

	// https://atomscan.com/validators/cosmosvaloper1crqm3598z6qmyn2kkcl9dz7uqs4qdqnr6s8jdn
	address = "cosmos1crqm3598z6qmyn2kkcl9dz7uqs4qdqnrlyn8pq"
)

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

	// List historical staking balances for the given address starting from the most recent after the given timestamp.
	stakesIter := stakingClient.Rewards.ListStakes(ctx, &api.ListStakesRequest{
		Parent:   rewards.Cosmos,
		PageSize: 20,
		Filter:   filter.WithAddress().Eq(address).String(),
	})

	// Iterates through the stakes and print them.
	for {
		stake, err := stakesIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing stakes: %s", err.Error())
		}

		marshaler := protojson.MarshalOptions{Indent: "\t"}
		marshaled, err := marshaler.Marshal(stake)
		if err != nil {
			log.Fatalf("error marshaling reward: %s", err.Error())
		}

		fmt.Println(string(marshaled))
	}
}
