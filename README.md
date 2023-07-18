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
2. Run `go build`
3. Download your API key from the Coinbase Cloud UI and save it as `.coinbase_cloud_api_key.json` at the root of this repo
4. Run `./staking-client-library-go`
5. You should see output like the following:
    ```
    2023/07/14 08:07:14 got protocol: protocols/polygon
    2023/07/14 08:07:14 got network: protocols/polygon/networks/mainnet
    2023/07/14 08:07:14 got action: protocols/polygon/networks/mainnet/actions/stake
    2023/07/14 08:07:14 got action: protocols/polygon/networks/mainnet/actions/unstake
    2023/07/14 08:07:14 got action: protocols/polygon/networks/mainnet/actions/restake
    2023/07/14 08:07:14 got action: protocols/polygon/networks/mainnet/actions/claim_rewards
    2023/07/14 08:07:14 got validator: protocols/polygon/networks/mainnet/validators/0x857679d69fe50e7b722f94acd2629d80c355163d
    2023/07/14 08:07:14 projectID, privateKey and delegatorAddress must be set
    ```
6. To proceed with the entire example.go, you will need to set the following variables in the example.go file and run the code again: 
   * `projectID` - this is the project ID of your Coinbase Cloud project
   * `privateKey` - this is the private key of the address with this you want to perform staking actions
   * `delegatorAddress` - this is the address with which you want to perform staking actions
