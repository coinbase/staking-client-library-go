/*
 * This example code demonstrates rewards client library usage for listing rewards
 */

package main

import (
	"bytes"
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
	"google.golang.org/protobuf/encoding/protojson"
)

/*
 * Run the code with 'go run examples/ethereum/list-rewards/main.go' to view the rewards for the first validator on the Ethereum network.
 * Or, to view rewards for any arbitrary validator, simply replace the public key below.
 */

const (
	// https://beaconcha.in/validator/1
	address = "0xa1d1ad0714035353258038e964ae9675dc0252ee22cea896825c01458e1807bfad2f9969338798548d9858a571f7425c"
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

	// Lists the rewards for the given address for the previous last 20 days, aggregated by day.
	rewardsIter := stakingClient.Rewards.ListRewards(ctx, &api.ListRewardsRequest{
		Parent:   rewards.Ethereum,
		PageSize: 200,
		Filter: filter.WithAddress().Eq(address).
			And(filter.WithPeriodEndTime().Gte(time.Now().AddDate(0, 0, -20))).
			And(filter.WithPeriodEndTime().Lt(time.Now())).String(),
	})

	// Iterates through the rewards and pretty print them.
	for {
		reward, err := rewardsIter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalf("error listing rewards: %s", err.Error())
		}

		marshaled, err := protojson.Marshal(reward)
		if err != nil {
			log.Fatalf("error marshaling reward: %s", err.Error())
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, marshaled, "", "\t"); err != nil {
			log.Fatalf("JSON parse error: %s", err)
			return
		}

		fmt.Println(prettyJSON.String())
	}
}
