/*
 * This example code demonstrates rewards client library usage for listing staking balances
 */

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
 * Run the code with 'go run examples/ethereum/list-stakes/main.go' to view the balances for the first validator on the Ethereum network.
 * Or, to view balances for any arbitrary validator, simply replace the public key below.
 */

const (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"

	// https://beaconcha.in/validator/1
	address = "0xa1d1ad0714035353258038e964ae9675dc0252ee22cea896825c01458e1807bfad2f9969338798548d9858a571f7425c"
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
		Parent:   rewards.Ethereum,
		PageSize: 200,
		Filter:   filter.WithAddress().Eq(address).String(),
	})

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
