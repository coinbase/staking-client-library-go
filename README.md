<img src='docs/images/banner.svg' width='1100' alt='Coinbase Staking API'>

# [Coinbase Staking API](https://github.com/coinbase/staking-client-library-go)

> Programmatic access to Coinbase's best-in-class staking infrastructure and services. :large_blue_circle:

[![Current version](https://img.shields.io/github/tag/coinbase/staking-client-library-go?color=3498DB&label=version)](https://github.com/coinbase/staking-client-library-go/releases) [![GitHub contributors](https://img.shields.io/github/contributors/coinbase/staking-client-library-go?color=3498DB)](https://github.com/coinbase/staking-client-library-go/graphs/contributors) [![GitHub Stars](https://img.shields.io/github/stars/coinbase/staking-client-library-go.svg?color=3498DB)](https://github.com/coinbase/staking-client-library-go/stargazers) [![GitHub](https://img.shields.io/github/license/coinbase/staking-client-library-go?color=3498DB)](https://github.com/coinbase/staking-client-library-go/blob/main/LICENSE)

## Overview

`staking-client-library-go` is the Go SDK for the **Coinbase Staking API** :large_blue_circle:.

The Coinbase Staking API empowers developers to deliver a fully-featured staking experience in their Web2 apps, wallets, or dApps using *one common interface* across protocols.

A traditional infrastructure-heavy staking integration can take months. Coinbase's Staking API enables onboarding within hours :sparkles:.

## Quick Start

1. Create and download an API key from the [Cloud Platform](https://portal.cloud.coinbase.com/access/api).
2. Place the key named `.coinbase_cloud_api_key.json` at the root of your repository.
3. Run one of the code samples [below](#stake-partial-eth-💠) or any of our [provided examples](./examples/) :rocket:.

### Stake Partial ETH :diamond_shape_with_a_dot_inside:

This code sample creates an ETH staking workflow. View the full code sample [here](examples/ethereum/create-workflow/main.go)

<details open>
  <summary>Code Sample</summary>

```golang
// examples/ethereum/create-workflow/main.go
package main

import (
    "context"
    "fmt"
    "log"

    "github.cbhq.net/cloud/staking-client-library-go/auth"
    "github.cbhq.net/cloud/staking-client-library-go/client"
    "github.cbhq.net/cloud/staking-client-library-go/client/options"
    stakingpb "github.cbhq.net/cloud/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

func main() {
    // TODO: Add your project ID found at cloud.coinbase.com or in your API key.
    projectID := ""

    ctx := context.Background()

    // Loads the API key from the default location.
    apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
    if err != nil {
        log.Fatalf("error loading API key: %s", err.Error())
    }

    // Creates the Coinbase Staking API client.
    stakingClient, err := client.New(ctx, options.WithAPIKey(apiKey))
    if err != nil {
        log.Fatalf("error instantiating staking client: %s", err.Error())
    }

    // Constructs the API request
    req := &stakingpb.CreateWorkflowRequest{
        Parent: fmt.Sprintf("projects/%s", projectID),
        Workflow: &stakingpb.Workflow{
            Action: "protocols/ethereum_kiln/networks/holesky/actions/stake",
            StakingParameters: &stakingpb.Workflow_EthereumKilnStakingParameters{
                EthereumKilnStakingParameters: &stakingpb.EthereumKilnStakingParameters{
                    Parameters: &stakingpb.EthereumKilnStakingParameters_StakeParameters{
                        StakeParameters: &stakingpb.EthereumKilnStakeParameters{
                            StakerAddress:             "0xdb816889F2a7362EF242E5a717dfD5B38Ae849FE",
                            IntegratorContractAddress: "0xA55416de5DE61A0AC1aa8970a280E04388B1dE4b",
                            Amount: &stakingpb.Amount{
                                Value:    "20",
                                Currency: "ETH",
                            },
                        },
                    },
                },
            },
            SkipBroadcast: true,
        },
    }

    workflow, err := stakingClient.Orchestration.CreateWorkflow(ctx, req)
    if err != nil {
        log.Fatalf("couldn't create workflow: %s", err.Error())
    }

    log.Printf("Workflow created: %s", workflow.Name)
}
```

</details>

   <details>
     <summary>Output</summary>

   ```text
   2024/03/28 11:43:49 Workflow created: projects/62376b2f-3f24-42c9-9025-d576a3c06d6f/workflows/ffbf9b45-c57b-49cb-a4d5-fdab66d8cb25
   ```

   </details>

### View Ethereum Rewards :moneybag:

This code sample returns rewards for an Ethereum validator address. View the full code sample [here](examples/ethereum/list-rewards/main.go).

<details open>
  <summary>Code Sample</summary>

```golang
// examples/ethereum/list-rewards/main.go
package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "time"

    "google.golang.org/api/iterator"

    "github.cbhq.net/cloud/staking-client-library-go/auth"
    "github.cbhq.net/cloud/staking-client-library-go/client"
    "github.cbhq.net/cloud/staking-client-library-go/client/options"
    rewardsV1 "github.cbhq.net/cloud/staking-client-library-go/client/rewards/v1"
    rewardspb "github.cbhq.net/cloud/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
)

func main() {
    ctx := context.Background()

    // Loads the API key from the default location.
    apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
    if err != nil {
        log.Fatalf("error loading API key: %s", err.Error())
    }

    // Creates the Coinbase Staking API client.
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
```

</details>

   <details>
     <summary>Output</summary>

   ```json
   {
      "address": "0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474",
      "period_identifier": {
         "date": "2024-03-26"
      },
      "aggregation_unit": 2,
      "period_start_time": {
         "seconds": 1711411200
      },
      "period_end_time": {
         "seconds": 1711497599
      },
      "total_earned_native_unit": {
         "amount": "0.00211503",
         "exp": "18",
         "ticker": "ETH",
         "raw_numeric": "2115030000000000"
      },
      "total_earned_usd": [
         {
            "source": 1,
            "conversion_time": {
               "seconds": 1711498140
            },
            "amount": {
               "amount": "7.58",
               "exp": "2",
               "ticker": "USD",
               "raw_numeric": "758"
            },
            "conversion_price": "3582.979980"
         }
      ],
      "protocol": "ethereum"
   }
   {
      "address": "0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474",
      "period_identifier": {
         "date": "2024-03-27"
      },
      "aggregation_unit": 2,
      "period_start_time": {
         "seconds": 1711497600
      },
      "period_end_time": {
         "seconds": 1711583999
      },
      "total_earned_native_unit": {
         "amount": "0.002034193",
         "exp": "18",
         "ticker": "ETH",
         "raw_numeric": "2034193000000000"
      },
      "total_earned_usd": [
         {
            "source": 1,
            "conversion_time": {
               "seconds": 1711584540
            },
            "amount": {
               "amount": "7.13",
               "exp": "2",
               "ticker": "USD",
               "raw_numeric": "713"
            },
            "conversion_price": "3504.580078"
         }
      ],
      "protocol": "ethereum"
   }
   ```

   </details>

## Documentation

There are numerous examples in the [`examples directory`](./examples) to help get you started. For even more, refer to our [documentation website](https://docs.cloud.coinbase.com/) for detailed definitions, API specifications, integration guides, and more!
