<img src='docs/images/banner.svg' width='1100' alt='Coinbase Staking API'>

# [Coinbase Staking API](https://github.com/coinbase/staking-client-library-go)

> Programmatic access to Coinbase's best-in-class staking infrastructure and services. :large_blue_circle:

[![Current version](https://img.shields.io/github/tag/coinbase/staking-client-library-go?color=3498DB&label=version)](https://github.com/coinbase/staking-client-library-go/releases) [![GitHub contributors](https://img.shields.io/github/contributors/coinbase/staking-client-library-go?color=3498DB)](https://github.com/coinbase/staking-client-library-go/graphs/contributors) [![GitHub Stars](https://img.shields.io/github/stars/coinbase/staking-client-library-go.svg?color=3498DB)](https://github.com/coinbase/staking-client-library-go/stargazers) [![GitHub](https://img.shields.io/github/license/coinbase/staking-client-library-go?color=3498DB)](https://github.com/coinbase/staking-client-library-go/blob/main/LICENSE)

## Overview

`staking-client-library-go` is the Go SDK for the **Coinbase Staking API** :large_blue_circle:.

The Coinbase Staking API empowers developers to deliver a fully-featured staking experience in their Web2 apps, wallets, or dApps using *one common interface* across protocols.

A traditional infrastructure-heavy staking integration can take months. Coinbase's Staking API enables onboarding within hours :sparkles:.

## Quick Start

Prerequisite: [Go 1.21+](https://go.dev/doc/install)

1. In a fresh directory, copy and paste one of the code samples below or any of our [provided examples](./examples) into an `example.go` file.
2. Create and download an API key from the [Cloud Platform](https://portal.cloud.coinbase.com/access/api).
3. Place the key named `.coinbase_cloud_api_key.json` at the root of your repository.
4. Setup a Go project and run the example :rocket:

   ```shell
   go mod init example
   go mod tidy
   go run example.go
   ```

### Stake Partial ETH :diamond_shape_with_a_dot_inside:

This code sample helps stake ETH via partial staking. View the full source [here](examples/ethereum/create-workflow/main.go)

<details open>

```golang
// examples/ethereum/create-workflow/main.go
package main

import (
   "context"
   "fmt"
   "log"

   "google.golang.org/protobuf/encoding/protojson"

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
                     StakerAddress: "0xdb816889F2a7362EF242E5a717dfD5B38Ae849FE",
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

   marshaled, err := protojson.MarshalOptions{Indent: "  ", Multiline: true}.Marshal(workflow)
   if err != nil {
      log.Fatalf("error marshaling reward: %s", err.Error())
   }

   fmt.Println(string(marshaled))
}
```

</details>

   <details>
     <summary>Output</summary>

   ```text
   {
     "name":  "workflows/357673b5-c3b7-4149-a887-a6119d32fbdd",
     "action":  "protocols/ethereum_kiln/networks/holesky/actions/stake",
     "ethereumKilnStakingParameters":  {
       "stakeParameters":  {
         "stakerAddress":  "0xdb816889F2a7362EF242E5a717dfD5B38Ae849FE",
         "integratorContractAddress":  "0xA55416de5DE61A0AC1aa8970a280E04388B1dE4b",
         "amount":  {
           "value":  "20",
           "currency":  "ETH"
         }
       }
     },
     "state":  "STATE_WAITING_FOR_EXT_BROADCAST",
     "steps":  [
       {
         "name":  "stake tx",
         "txStepOutput":  {
           "unsignedTx":  "02f3824268068502540be4008503743b80ba83061a8094a55416de5de61a0ac1aa8970a280e04388b1de4b14843a4b66f1c0808080",
           "state":  "STATE_PENDING_EXT_BROADCAST"
         }
       }
     ],
     "createTime":  "2024-05-08T17:30:56.633391257Z",
     "updateTime":  "2024-05-08T17:30:56.633391257Z"
   }
   ```

   </details>

### Stake SOL :diamond_shape_with_a_dot_inside:

This code sample helps stake SOL from a given user wallet. View the full source [here](examples/solana/create-workflow/main.go)

<details open>

```golang
// examples/solana/create-workflow/main.go
package main

import (
    "context"
    "fmt"
    "log"

    "google.golang.org/protobuf/encoding/protojson"

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
                            WalletAddress: "8rMGARtkJY5QygP1mgvBFLsE9JrvXByARJiyNfcSE5Z",
                            Amount: &api.Amount{
                                Value:    "100000000",
                                Currency: "SOL",
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

    marshaled, err := protojson.MarshalOptions{Indent: "  ", Multiline: true}.Marshal(workflow)
    if err != nil {
       log.Fatalf("error marshaling reward: %s", err.Error())
    }

    fmt.Println(string(marshaled))
}
```

</details>

   <details>
     <summary>Output</summary>

   ```text
   {
     "name":  "workflows/f68fc083-18bd-48a0-a10a-7c673e9ec5b7",
     "action":  "protocols/solana/networks/devnet/actions/stake",
     "solanaStakingParameters":  {
       "stakeParameters":  {
         "walletAddress":  "8rMGARtkJY5QygP1mgvBFLsE9JrvXByARJiyNfcSE5Z",
         "validatorAddress":  "GkqYQysEGmuL6V2AJoNnWZUz2ZBGWhzQXsJiXm2CLKAN",
         "amount":  {
           "value":  "100000000",
           "currency":  "SOL"
         },
         "priorityFee":  {}
       }
     },
     "state":  "STATE_WAITING_FOR_EXT_BROADCAST",
     "steps":  [
       {
         "name":  "stake tx",
         "txStepOutput":  {
           "unsignedTx":  "66hEYYWnwGWkGpMKF2H2sCzxnmoAfY8LPnYMgWdY6rC7hX2H6DEE2YdPxECFx8FeeNmea8N87L4KuZ6dirYXZi9XNr5uPJdf8W1jdShcSwzSmmqz4SA7dmFjdTM19hNEu7hMMF7C2RtjZj4qCRvArcnyjj76r5hJrm1o1RozjjZCyvgNqDGHYoeej9MPwoMUEaY6h2iKBh1hnkYFCA1tyXEP8xX3f1jbnae8jzW2Zkc62GDw2gKWusQ3KtRz3wRLdqWT9tbhEk6Hekqbw4sPXSPevsiYHPVX9mQJRdNkoYovBRXv3KQaQ7dv6isgyax7S53yoMRgCvfuhYxk9WhzR4fkAxYB26qeqpdUJrgvSpaw4T3iNBYsG7KZzvGUg4NWG1BaBDuvnG1x7YL3gyJd5QMWQ6jq6yuGgupjNn7zP7EcxtbvpP2EVfrFnzmX4LEgQh4MxshMFpNas2tQXQd12Vv9vq4nZt2BEr2Jh67Q9vnKF22td1XaAL1MvsTmvWWSKviyZkZQzTXsqUGFtox1f8Unwj75sNCQWYUh4PHmiUjWGmQVhyQKbEqG6PeqDyy9YTopamSD2ajDrhak5fsnczdXo166cjnzQAJZW7tN2T6jHJy2KNmDdL16qPR4HqKKXpWquf1NuTPuJ7ikfmJxWBp7gHrMF3z5P84hp4xT4V2D4eHGLMTWDhs4cQghVXRvynPUeUDSf3TWzHfYwVEDFNpFNhX62FP7aJBVp28R8nHTt7riymgkw9LdjhPMxQPoRW3hCG4UcQ9kJ7Aywcij9SVcbfaESoEz7anV1j6HFrXiQsgiSbeCj4iiXtYy9aDbVjiuv2v31kvoE6kb1s7osVoeK1mn7AxkPozMxTVhwca9BTMuHTYpFP1QTVAwsYSCXa6KuoXEgLKZn1c63ijQXXGENLjd17JqV3FK8x2Vurkunws8pAb",
           "state":  "STATE_PENDING_EXT_BROADCAST"
         }
       }
     ],
     "createTime":  "2024-05-08T17:33:40.843044346Z",
     "updateTime":  "2024-05-08T17:33:40.843044346Z"
   }
   ```

   </details>

### View Ethereum Rewards :moneybag:

This code sample helps view rewards for an Ethereum validator address. View the full source [here](examples/ethereum/list-rewards/main.go).

<details open>

```golang
// examples/ethereum/list-rewards/main.go
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

        marshaled, err := protojson.MarshalOptions{Indent: "  ", Multiline: true}.Marshal(reward)
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

There are numerous examples in the [`examples directory`](./examples) to help get you started. For even more, refer to our [documentation website](https://docs.cdp.coinbase.com/staking/docs/welcome) for detailed definitions, API specifications, integration guides, and more!

## Contributing

Thanks for considering contributing to the project! Please refer to [our contribution guide](./CONTRIBUTING.md).

## Contact Us

If you have any questions, please reach out to us in the #staking channel on our [Discord](https://discord.com/invite/cdp) server.
