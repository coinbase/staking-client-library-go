/*
 * This example code, demonstrates rewards client library usage for listing rewards
 */

package main

import (
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	rewardsV1 "github.com/coinbase/staking-client-library-go/client/rewards/v1"
	rewardspb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
)

const (
	// TODO: Replace address as per your requirement.
	address = "cosmos1crqm3598z6qmyn2kkcl9dz7uqs4qdqnrlyn8pq"
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

	// List all rewards for the given address, aggregated by day, for epochs that ended in the last 30 days.

	now := time.Now()
	nowDateOnly := now.Truncate(24 * time.Hour)

	thirtyDaysAgo := now.AddDate(0, 0, -30)
	thirtyDaysAgoDateOnly := thirtyDaysAgo.Truncate(24 * time.Hour)

	rewardsIter := stakingClient.Rewards.ListRewards(ctx, &rewardspb.ListRewardsRequest{
		Parent:   rewardspb.ProtocolResourceName{Protocol: "cosmos"}.String(),
		PageSize: 20,
		Filter: rewardsV1.WithAddress().Eq(address).
			And(rewardsV1.WithDate().Gte(thirtyDaysAgoDateOnly.Format("2006-01-02"))).
			And(rewardsV1.WithDate().Lt(nowDateOnly.Format("2006-01-02"))).String(),
	})

	count := 0
	for {
		reward, err := rewardsIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing rewards: %s", err.Error())
		}

		d, err := protojson.Marshal(reward)
		if err != nil {
			log.Fatalf("error marshalling reward object: %s", err.Error())
		}

		log.Printf("[%d] Reward details: %s", count, d)
		count += 1
	}
}
