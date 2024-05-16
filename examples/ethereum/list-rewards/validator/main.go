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

	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	"github.com/coinbase/staking-client-library-go/client/rewards"
	filter "github.com/coinbase/staking-client-library-go/client/rewards/rewardsfilter"
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
)
 
 const (
	apiKeyName    = "your-api-key-name"
	apiPrivateKey = "your-api-private-key"
	
	 // https://beaconcha.in/validator/1
	 validatorAddress = "0xa1d1ad0714035353258038e964ae9675dc0252ee22cea896825c01458e1807bfad2f9969338798548d9858a571f7425c"
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
 
	  // Lists the rewards for the given validator address between February 25, 2024 and February 27th, 2024 aggregated by day.
	  validatorRewardsIter := stakingClient.Rewards.ListRewards(ctx, &api.ListRewardsRequest{
        Parent:   rewards.Ethereum,
        PageSize: 200,
        Filter: filter.WithAddress().Eq(validatorAddress).
		And(filter.WithPeriodEndTime().Gte(time.Date(2024, 2, 25, 0, 0, 0, 0, time.Local))).
		And(filter.WithPeriodEndTime().Lt(time.Date(2024, 2, 27, 0, 0, 0, 0, time.Local))).String(),
    })

	 // Iterate through the validator rewards and pretty print them.
	 for {
		reward, err := validatorRewardsIter.Next()
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
