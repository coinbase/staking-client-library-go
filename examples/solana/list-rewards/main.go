/*
 * This example code list historical staking rewards of a validator on the Solana blockchain
 * A user can view historical staking rewards for any validator on the Solana blockchain by replacing the address below
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	"github.com/coinbase/staking-client-library-go/client/rewards"
	filter "github.com/coinbase/staking-client-library-go/client/rewards/rewardsfilter"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

/*
 * Run the code with 'go run examples/solana/list-rewards/main.go' to view the rewards for Coinbase Cloud's public validator.
 * Or, to view rewards for any arbitrary validator, simply replace the address below with any validator on the Solana blockchain.
 */

const (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"

	// https://solanabeach.io/validator/6D2jqw9hyVCpppZexquxa74Fn33rJzzBx38T58VucHx9
	address = "6D2jqw9hyVCpppZexquxa74Fn33rJzzBx38T58VucHx9"
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

	// List all rewards for the given address, aggregated by epoch, for epochs that ended in the last 30 days.
	rewardsIter := stakingClient.Rewards.ListRewards(ctx, &api.ListRewardsRequest{
		Parent:   rewards.Solana,
		PageSize: 20,
		Filter: filter.WithAddress().Eq(address).
			And(filter.WithPeriodEndTime().Gte(time.Now().AddDate(0, 0, -30))).
			And(filter.WithPeriodEndTime().Lt(time.Now())).String(),
	})

	// Iterates through the rewards and print them.
	for {
		reward, err := rewardsIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing rewards: %s", err.Error())
		}

		marshaler := protojson.MarshalOptions{Indent: "\t"}
		marshaled, err := marshaler.Marshal(reward)
		if err != nil {
			log.Fatalf("error marshaling reward: %s", err.Error())
		}

		fmt.Println(string(marshaled))
	}
}
