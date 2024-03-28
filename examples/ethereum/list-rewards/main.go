/*
 * This example code demonstrates rewards client library usage for listing rewards
 */

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/iterator"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	rewardsV1 "github.com/coinbase/staking-client-library-go/client/rewards/v1"
	rewardspb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
)

func main() {
	ctx := context.Background()

	// Loads the API key from the default location.
	apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	// Creates the Coinbase Staking API client
	stakingClient, err := client.New(ctx, options.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	// Lists the rewards for the given address for the previous last 2 days, aggregated by day.
	rewardsIter := stakingClient.Rewards.ListRewards(ctx, &rewardspb.ListRewardsRequest{
		Parent:   rewardspb.ProtocolResourceName{Protocol: "ethereum"}.String(),
		PageSize: 200,
		Filter: rewardsV1.WithAddress().Eq("0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474").
			And(rewardsV1.WithPeriodEndTime().Gte(time.Now().AddDate(0, 0, -2))).
			And(rewardsV1.WithPeriodEndTime().Lt(time.Now())).String(),
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

		marshaled, err := json.MarshalIndent(reward, "", "   ")
		if err != nil {
			log.Fatalf("error marshaling reward: %s", err.Error())
		}

		fmt.Printf(string(marshaled))
	}
}
