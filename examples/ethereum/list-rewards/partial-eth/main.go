/*
 * This example code demonstrates rewards client library usage for listing rewards
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

const (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"

	partialETHAddress = "0x60c7e246344ae3856cf9abe3a2e258d495fc39e0"
)

func main() {
	ctx := context.Background()

	// Loads the API key.
	apiKey, err := auth.NewAPIKey(auth.WithAPIKeyName(apiKeyName, apiPrivateKey))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	// Creates the Coinbase Staking API client
	stakingClient, err := client.New(ctx, options.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	// Lists the rewards for the given partial eth address between May 1st, 2024 and May 3rd, 2024 aggregated by day.
	partialETHRewardsIter := stakingClient.Rewards.ListRewards(ctx, &api.ListRewardsRequest{
		Parent:   rewards.Ethereum,
		PageSize: 200,
		Filter: filter.WithAddress().Eq(partialETHAddress).
			And(filter.WithPeriodEndTime().Gte(time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local))).
			And(filter.WithPeriodEndTime().Lt(time.Date(2024, 5, 3, 0, 0, 0, 0, time.Local))).String(),
	})

	// Iterate through the partial eth rewards and pretty print them.
	for {
		reward, err := partialETHRewardsIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing rewards: %s", err.Error())
		}

		marshaled, err := protojson.MarshalOptions{Indent: "  ", Multiline: true}.Marshal(reward)
		if err != nil {
			log.Fatalf("error marshaling reward: %s", err.Error())
		}

		fmt.Println(string(marshaled))
	}
}
