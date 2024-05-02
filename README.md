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
    "log"

    "github.com/coinbase/staking-client-library-go/auth"
    "github.com/coinbase/staking-client-library-go/client"
    "github.com/coinbase/staking-client-library-go/client/options"
    api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
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

    req := &api.CreateWorkflowRequest{
        Workflow: &api.Workflow{
            Action: "protocols/ethereum_kiln/networks/holesky/actions/stake",
            StakingParameters: &api.Workflow_EthereumKilnStakingParameters{
                EthereumKilnStakingParameters: &api.EthereumKilnStakingParameters{
                    Parameters: &api.EthereumKilnStakingParameters_StakeParameters{
                        StakeParameters: &api.EthereumKilnStakeParameters{
                            StakerAddress:             "0xdb816889F2a7362EF242E5a717dfD5B38Ae849FE",
                            Amount: &api.Amount{
                                Value:    "20",
                                Currency: "ETH",
                            },
                        },
                    },
                },
            },
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
   2024/03/28 11:43:49 Workflow created: workflows/ffbf9b45-c57b-49cb-a4d5-fdab66d8cb25
   ```

   </details>

### Stake SOL :diamond_shape_with_a_dot_inside:

This code sample creates an SOL staking workflow. View the full code sample [here](examples/solana/create-workflow/main.go)

<details open>
  <summary>Code Sample</summary>

```golang
// examples/solana/create-workflow/main.go
package main

import (
    "context"
    "log"

    "github.com/coinbase/staking-client-library-go/auth"
    "github.com/coinbase/staking-client-library-go/client"
    "github.com/coinbase/staking-client-library-go/client/options"
    api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
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

	req := &api.CreateWorkflowRequest{
		Workflow: &api.Workflow{
			Action: "protocols/solana/networks/devnet/actions/stake",
			StakingParameters: &api.Workflow_SolanaStakingParameters{
				SolanaStakingParameters: &api.SolanaStakingParameters{
					Parameters: &api.SolanaStakingParameters_StakeParameters{
						StakeParameters: &api.SolanaStakeParameters{
							WalletAddress: walletAddress,
							Amount: &api.Amount{
								Value:    amount,
								Currency: currency,
							},
						},
					},
				},
			},
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
   2024/03/28 11:43:49 Workflow created: workflows/6bd9fd82-8b9d-4a49-9039-f95bb850a7a2
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

        marshaler := protojson.MarshalOptions{Indent: "\t"}
        marshaled, err := marshaler.Marshal(reward)
        if err != nil {
            log.Fatalf("error marshaling reward: %s", err.Error())
        }

        fmt.Println(string(marshaled))
    }
}
```

</details>

   <details>
     <summary>Output</summary>

   ```json
   {
        "address": "0xa1d1ad0714035353258038e964ae9675dc0252ee22cea896825c01458e1807bfad2f9969338798548d9858a571f7425c",
        "date": "2024-04-20",
        "aggregationUnit": "DAY",
        "periodStartTime": "2024-04-20T00:00:00Z",
        "periodEndTime": "2024-04-20T23:59:59Z",
        "totalEarnedNativeUnit": {
            "amount": "0.002118354",
            "exp": "18",
            "ticker": "ETH",
            "rawNumeric": "2118354000000000"
        },
        "totalEarnedUsd": [
            {
                "source": "COINBASE_EXCHANGE",
                "conversionTime": "2024-04-21T00:09:00Z",
                "amount": {
                    "amount": "6.67",
                    "exp": "2",
                    "ticker": "USD",
                    "rawNumeric": "667"
                },
                "conversionPrice": "3145.550049"
            }
        ],
        "protocol": "ethereum"
    }
    {
        "address": "0xa1d1ad0714035353258038e964ae9675dc0252ee22cea896825c01458e1807bfad2f9969338798548d9858a571f7425c",
        "date": "2024-04-21",
        "aggregationUnit": "DAY",
        "periodStartTime": "2024-04-21T00:00:00Z",
        "periodEndTime": "2024-04-21T23:59:59Z",
        "totalEarnedNativeUnit": {
            "amount": "0.00211564",
            "exp": "18",
            "ticker": "ETH",
            "rawNumeric": "2115640000000000"
        },
        "totalEarnedUsd": [
            {
                "source": "COINBASE_EXCHANGE",
                "conversionTime": "2024-04-22T00:09:00Z",
                "amount": {
                    "amount": "6.68",
                    "exp": "2",
                    "ticker": "USD",
                    "rawNumeric": "668"
                },
                "conversionPrice": "3155.449951"
            }
        ],
        "protocol": "ethereum"
    }
   ```

   </details>

## Documentation

There are numerous examples in the [`examples directory`](./examples) to help get you started. For even more, refer to our [documentation website](https://docs.cloud.coinbase.com/) for detailed definitions, API specifications, integration guides, and more!
