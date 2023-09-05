# Staking API Go Client Library

This repository contains the Protocol Buffer definitions for the Coinbase **Staking API**, as well as the Go client libraries generated from them.

## Overview

Staking API provides a set of APIs to aid in non-custodial staking for multiple protocols and networks.

# Documentation

In order to self-host the documentation run please use the swagger-ui docker container like so:

```bash
docker run -p 8080:8080 -e SWAGGER_JSON=/doc/coinbase/staking/v1alpha1/api.swagger.json -v $(PWD)/doc/openapi:/doc swaggerapi/swagger-ui
```

## Prerequisites

- [Golang 1.19+](https://go.dev/learn/)

## Repository Structure
- [`auth/`](./auth/) contains the authentication-related code for accessing Coinbase Cloud APIs.
- [`client/`](./client/) contains client instantiation helpers for Staking APIs.
- [`gen/`](./gen/) contains Go code generated from the Protocol Buffers.
- [`protos/`](./protos/) contains the Protocol Buffers that define the Staking APIs.

## Module Installation
```
go get github.com/coinbase/staking-client-library-go
```

## Get Started
To test that your API Key gives you access as expected to the Staking APIs:

1. Clone this GitHub repo
2. Run `go build -o staking-client-example examples/example.go`
3. Download your API key from the Coinbase Cloud UI and save it as `.coinbase_cloud_api_key.json` at the root of this repo
4. Run `./staking-client-example`
5. You should see output like the following:
    ```
   2023/09/01 17:52:04 got protocol: protocols/polygon
   2023/09/01 17:52:04 got protocol: protocols/solana
   2023/09/01 17:52:04 got network: protocols/polygon/networks/mainnet
   2023/09/01 17:52:04 got action: protocols/polygon/networks/mainnet/actions/stake
   2023/09/01 17:52:04 got action: protocols/polygon/networks/mainnet/actions/unstake
   2023/09/01 17:52:04 got action: protocols/polygon/networks/mainnet/actions/restake
   2023/09/01 17:52:04 got action: protocols/polygon/networks/mainnet/actions/claim_rewards
   2023/09/01 17:52:04 got validator: protocols/polygon/networks/mainnet/validators/0x857679d69fE50E7B722f94aCd2629d80C355163d
    ```
### Create an Ethereum Kiln workflow
To test creating an Ethereum Kiln workflow perform the following:

1. Run `go build -o staking-client-kiln-example examples/ethereum_kiln/example.go`
2. Run `./staking-client-kiln-example`
3. To proceed with the entire example.go, you will need to set the following variables in the example.go file and run the code again:
   * `projectID` - this is the project ID of your Coinbase Cloud project
   * `privateKey` - this is the private key of the address with this you want to perform staking actions
   * `stakerAddress` - this is the address with which you want to perform staking actions

### Create a Polygon workflow
To test creating a Polygon workflow perform the following:

1. Run `go build -o staking-client-polygon-example examples/polygon/example.go`
2. Run `./staking-client-polygon-example`
3. To proceed with the entire example.go, you will need to set the following variables in the example.go file and run the code again:
   * `projectID` - this is the project ID of your Coinbase Cloud project
   * `privateKey` - this is the private key of the address with this you want to perform staking actions
   * `delegatorAddress` - this is the address with which you want to perform staking actions
