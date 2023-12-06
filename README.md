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
2. Download your API key from the Coinbase Cloud UI and save it as `.coinbase_cloud_api_key.json` at the root of this repo
3. Run `go run examples/example.go`
4. You should see output like the following:
    ```
   2023/09/25 15:01:03 got protocol: protocols/ethereum_kiln
   2023/09/25 15:01:04 got network: protocols/ethereum_kiln/networks/mainnet
   2023/09/25 15:01:04 got action: protocols/ethereum_kiln/networks/goerli/actions/stake
   2023/09/25 15:01:04 got action: protocols/ethereum_kiln/networks/goerli/actions/unstake
   2023/09/25 15:01:04 got action: protocols/ethereum_kiln/networks/goerli/actions/claim_rewards
    ```
### Create an Ethereum Kiln workflow
To test creating an Ethereum Kiln workflow perform the following:

1. Open `examples/ethereum_kiln/example.go` and set the following variables:
   * `projectID` - this is the project ID of your Coinbase Cloud project
   * `privateKey` - this is the private key of the address with this you want to perform staking actions
   * `stakerAddress` - this is the address with which you want to perform staking actions
   * `integratorContractAddress` - this is the integrator contract address to which you want to stake. This is typically procured by engaging with the Coinbase Cloud sales team.
2. Run `go run examples/ethereum_kiln/example.go`

### Create a Polygon workflow
To test creating a Polygon workflow perform the following:

1. Open `examples/polygon/example.go` and set the following variables:
   * `projectID` - this is the project ID of your Coinbase Cloud project
   * `privateKey` - this is the private key of the address with this you want to perform staking actions
   * `delegatorAddress` - this is the address with which you want to perform staking actions
2. Run `go run examples/polygon/example.go`

### Listing workflows
To test listing workflows within a project, perform the following:

1. Open `examples/example_list_workflows.go` and set the following variables:
   * `projectID` - this is the project ID of your Coinbase Cloud project
2. Run `go run examples/example_list_workflows.go`
