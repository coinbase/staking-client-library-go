package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	"github.com/coinbase/staking-client-library-go/client/rewards"
	filter "github.com/coinbase/staking-client-library-go/client/rewards/rewards_filter"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
	"google.golang.org/api/iterator"
)

/*
 * Run the code with 'go run main.go' to view the rewards for Coinbase Cloud's public validator.
 * Or, to view rewards for any arbitrary validator, simply replace the address below with any validator on the Cosmos blockchain.
 */

const (
	// https://atomscan.com/validators/cosmosvaloper1crqm3598z6qmyn2kkcl9dz7uqs4qdqnr6s8jdn
	address = "cosmos1crqm3598z6qmyn2kkcl9dz7uqs4qdqnrlyn8pq"
)

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
	rewardsIter := stakingClient.Rewards.ListRewards(ctx, &api.ListRewardsRequest{
		Parent:   rewards.Cosmos,
		PageSize: 20,
		Filter: filter.WithAddress().Eq(address).
			And(filter.WithPeriodEndTime().Gte(time.Now().AddDate(0, 0, -20))).
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

		marshaled, err := json.MarshalIndent(reward, "", "   ")
		if err != nil {
			log.Fatalf("error marshaling reward: %s", err.Error())
		}

		fmt.Println(string(marshaled))
	}
}
