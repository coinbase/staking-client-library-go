/*
 * This example code demonstrates rewards client library usage for listing staking balances
 */

package main

import (
	"context"
	"errors"
	"log"

	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	rewardspb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
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

	// List staking activities for a given protocol.
	stakesIter := stakingClient.Rewards.ListStakes(ctx, &rewardspb.ListStakesRequest{
		Parent:   rewardspb.ProtocolResourceName{Protocol: "ethereum"}.String(),
		PageSize: 200,
		Filter:   `address='0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474' AND evaluation_time>='2023-12-12T07:25:11-04:00' AND evaluation_time<='2023-12-12T08:20:50-04:00'`,
	})

	count := 0
	for {
		stake, err := stakesIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing stakes: %s", err.Error())
		}

		d, err := protojson.Marshal(stake)
		if err != nil {
			log.Fatalf("error marshalling stake object: %s", err.Error())
		}

		log.Printf("[%d] Stake details: %s", count, d)
		count += 1
	}
}
