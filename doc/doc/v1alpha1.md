# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [coinbase/staking/v1alpha1/action.proto](#coinbase_staking_v1alpha1_action-proto)
    - [Action](#coinbase-staking-v1alpha1-Action)
    - [ListActionsRequest](#coinbase-staking-v1alpha1-ListActionsRequest)
    - [ListActionsResponse](#coinbase-staking-v1alpha1-ListActionsResponse)
  
- [coinbase/staking/v1alpha1/protocol.proto](#coinbase_staking_v1alpha1_protocol-proto)
    - [ListProtocolsRequest](#coinbase-staking-v1alpha1-ListProtocolsRequest)
    - [ListProtocolsResponse](#coinbase-staking-v1alpha1-ListProtocolsResponse)
    - [Protocol](#coinbase-staking-v1alpha1-Protocol)
  
- [coinbase/staking/v1alpha1/network.proto](#coinbase_staking_v1alpha1_network-proto)
    - [ListNetworksRequest](#coinbase-staking-v1alpha1-ListNetworksRequest)
    - [ListNetworksResponse](#coinbase-staking-v1alpha1-ListNetworksResponse)
    - [Network](#coinbase-staking-v1alpha1-Network)
  
- [coinbase/staking/v1alpha1/validator.proto](#coinbase_staking_v1alpha1_validator-proto)
    - [ListValidatorsRequest](#coinbase-staking-v1alpha1-ListValidatorsRequest)
    - [ListValidatorsResponse](#coinbase-staking-v1alpha1-ListValidatorsResponse)
    - [Validator](#coinbase-staking-v1alpha1-Validator)
  
- [coinbase/staking/v1alpha1/common.proto](#coinbase_staking_v1alpha1_common-proto)
    - [Amount](#coinbase-staking-v1alpha1-Amount)
  
- [coinbase/staking/v1alpha1/ethereum_kiln.proto](#coinbase_staking_v1alpha1_ethereum_kiln-proto)
    - [EthereumKilnClaimRewardsParameters](#coinbase-staking-v1alpha1-EthereumKilnClaimRewardsParameters)
    - [EthereumKilnStakeParameters](#coinbase-staking-v1alpha1-EthereumKilnStakeParameters)
    - [EthereumKilnStakingContextDetails](#coinbase-staking-v1alpha1-EthereumKilnStakingContextDetails)
    - [EthereumKilnStakingContextParameters](#coinbase-staking-v1alpha1-EthereumKilnStakingContextParameters)
    - [EthereumKilnStakingParameters](#coinbase-staking-v1alpha1-EthereumKilnStakingParameters)
    - [EthereumKilnUnstakeParameters](#coinbase-staking-v1alpha1-EthereumKilnUnstakeParameters)
  
- [coinbase/staking/v1alpha1/polygon.proto](#coinbase_staking_v1alpha1_polygon-proto)
    - [PolygonClaimRewardsParameters](#coinbase-staking-v1alpha1-PolygonClaimRewardsParameters)
    - [PolygonRestakeParameters](#coinbase-staking-v1alpha1-PolygonRestakeParameters)
    - [PolygonStakeParameters](#coinbase-staking-v1alpha1-PolygonStakeParameters)
    - [PolygonStakingParameters](#coinbase-staking-v1alpha1-PolygonStakingParameters)
    - [PolygonUnstakeParameters](#coinbase-staking-v1alpha1-PolygonUnstakeParameters)
  
- [coinbase/staking/v1alpha1/solana.proto](#coinbase_staking_v1alpha1_solana-proto)
    - [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions)
    - [SolanaCreateStakeAccountParameters](#coinbase-staking-v1alpha1-SolanaCreateStakeAccountParameters)
    - [SolanaDeactivateStakeParameters](#coinbase-staking-v1alpha1-SolanaDeactivateStakeParameters)
    - [SolanaDelegateStakeParameters](#coinbase-staking-v1alpha1-SolanaDelegateStakeParameters)
    - [SolanaMergeStakeParameters](#coinbase-staking-v1alpha1-SolanaMergeStakeParameters)
    - [SolanaSplitStakeParameters](#coinbase-staking-v1alpha1-SolanaSplitStakeParameters)
    - [SolanaStakingParameters](#coinbase-staking-v1alpha1-SolanaStakingParameters)
    - [SolanaWithdrawStakeParameters](#coinbase-staking-v1alpha1-SolanaWithdrawStakeParameters)
  
- [coinbase/staking/v1alpha1/workflow.proto](#coinbase_staking_v1alpha1_workflow-proto)
    - [CreateWorkflowRequest](#coinbase-staking-v1alpha1-CreateWorkflowRequest)
    - [GetWorkflowRequest](#coinbase-staking-v1alpha1-GetWorkflowRequest)
    - [ListWorkflowsRequest](#coinbase-staking-v1alpha1-ListWorkflowsRequest)
    - [ListWorkflowsResponse](#coinbase-staking-v1alpha1-ListWorkflowsResponse)
    - [PerformWorkflowStepRequest](#coinbase-staking-v1alpha1-PerformWorkflowStepRequest)
    - [RefreshWorkflowStepRequest](#coinbase-staking-v1alpha1-RefreshWorkflowStepRequest)
    - [TxStepOutput](#coinbase-staking-v1alpha1-TxStepOutput)
    - [WaitStepOutput](#coinbase-staking-v1alpha1-WaitStepOutput)
    - [Workflow](#coinbase-staking-v1alpha1-Workflow)
    - [WorkflowStep](#coinbase-staking-v1alpha1-WorkflowStep)
  
    - [TxStepOutput.State](#coinbase-staking-v1alpha1-TxStepOutput-State)
    - [WaitStepOutput.State](#coinbase-staking-v1alpha1-WaitStepOutput-State)
    - [WaitStepOutput.WaitUnit](#coinbase-staking-v1alpha1-WaitStepOutput-WaitUnit)
    - [Workflow.State](#coinbase-staking-v1alpha1-Workflow-State)
  
- [coinbase/staking/v1alpha1/staking_context.proto](#coinbase_staking_v1alpha1_staking_context-proto)
    - [ViewStakingContextRequest](#coinbase-staking-v1alpha1-ViewStakingContextRequest)
    - [ViewStakingContextResponse](#coinbase-staking-v1alpha1-ViewStakingContextResponse)
  
- [coinbase/staking/v1alpha1/api.proto](#coinbase_staking_v1alpha1_api-proto)
    - [StakingService](#coinbase-staking-v1alpha1-StakingService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="coinbase_staking_v1alpha1_action-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/action.proto



<a name="coinbase-staking-v1alpha1-Action"></a>

### Action
An Action resource, which represents an action you may take on a network,
posted to a validator. (i.e. stake, unstake)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the Action. Format: protocols/{protocolName}/networks/{networkName}/actions/{actionName} Ex: protocols/polygon/networks/goerli/validators/stake |






<a name="coinbase-staking-v1alpha1-ListActionsRequest"></a>

### ListActionsRequest
The request message for ListActions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parent | [string](#string) |  | The resource name of the parent that owns the collection of actions. Format: protocols/{protocol}/networks/{network} |






<a name="coinbase-staking-v1alpha1-ListActionsResponse"></a>

### ListActionsResponse
The response message for ListActions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [Action](#coinbase-staking-v1alpha1-Action) | repeated | The list of actions. |





 

 

 

 



<a name="coinbase_staking_v1alpha1_protocol-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/protocol.proto



<a name="coinbase-staking-v1alpha1-ListProtocolsRequest"></a>

### ListProtocolsRequest
The request message for ListProtocols.






<a name="coinbase-staking-v1alpha1-ListProtocolsResponse"></a>

### ListProtocolsResponse
The response message for ListProtocols.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protocols | [Protocol](#coinbase-staking-v1alpha1-Protocol) | repeated | The list of protocols. |






<a name="coinbase-staking-v1alpha1-Protocol"></a>

### Protocol
A Protocol resource (i.e. polygon, ethereum, etc.)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the Protocol. Format: protocols/{protocolName} Ex: protocols/polygon |





 

 

 

 



<a name="coinbase_staking_v1alpha1_network-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/network.proto



<a name="coinbase-staking-v1alpha1-ListNetworksRequest"></a>

### ListNetworksRequest
The request message for ListNetworks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parent | [string](#string) |  | The resource name of the parent that owns the collection of networks. Format: protocols/{protocol} |






<a name="coinbase-staking-v1alpha1-ListNetworksResponse"></a>

### ListNetworksResponse
The response message for ListNetworks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| networks | [Network](#coinbase-staking-v1alpha1-Network) | repeated | The list of networks. |






<a name="coinbase-staking-v1alpha1-Network"></a>

### Network
A Network resource, which represents a blockchain network.
(i.e. mainnet, testnet, etc.)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the Network. Format: protocols/{protocolName}/networks/{networkName} Ex: protocols/polygon/networks/goerli |
| is_mainnet | [bool](#bool) |  | Represents if the network is the mainnet network for the given protocol. |





 

 

 

 



<a name="coinbase_staking_v1alpha1_validator-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/validator.proto



<a name="coinbase-staking-v1alpha1-ListValidatorsRequest"></a>

### ListValidatorsRequest
The request message for ListValidators.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parent | [string](#string) |  | The resource name of the parent that owns the collection of validators. Format: protocols/{protocol}/networks/{network} |
| page_size | [int32](#int32) |  | The maximum number of validators to return. The service may return fewer than this value.

If unspecified, 50 validators will be returned. The maximum value is 1000; values over 1000 will be floored to 1000. |
| page_token | [string](#string) |  | A page token as part of the response of a previous call. Provide this to retrieve the next page.

When paginating, all other parameters must match the previous request to list resources. |






<a name="coinbase-staking-v1alpha1-ListValidatorsResponse"></a>

### ListValidatorsResponse
The response message for ListValidators.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validators | [Validator](#coinbase-staking-v1alpha1-Validator) | repeated | The list of validators. |
| next_page_token | [string](#string) |  | A token which can be provided as `page_token` to retrieve the next page. If this field is omitted, there are no additional pages. |






<a name="coinbase-staking-v1alpha1-Validator"></a>

### Validator
A Validator resource, which represents an active validator
for the given protocol network which you can submit an action
to.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the Validator. Format: protocols/{protocolName}/networks/{networkName}/validators/{validatorName} Ex: protocols/polygon/networks/goerli/validators/0x857679d69fE50E7B722f94aCd2629d80C355163d |
| validator_address | [string](#string) |  | The public address of the validator that you may perform workflow actions on. |
| commission_rate | [float](#float) |  | The rate of commission for the validator |





 

 

 

 



<a name="coinbase_staking_v1alpha1_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/common.proto



<a name="coinbase-staking-v1alpha1-Amount"></a>

### Amount
The amount of a token you wish to perform an action
with.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | The total value of the token. |
| currency | [string](#string) |  | The name of the token. |





 

 

 

 



<a name="coinbase_staking_v1alpha1_ethereum_kiln-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/ethereum_kiln.proto



<a name="coinbase-staking-v1alpha1-EthereumKilnClaimRewardsParameters"></a>

### EthereumKilnClaimRewardsParameters
The parameters required for the claim rewards action on Ethereum Kiln.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staker_address | [string](#string) |  | The address you wish to claim rewards for. |
| integrator_contract_address | [string](#string) |  | The address of the integrator contract |






<a name="coinbase-staking-v1alpha1-EthereumKilnStakeParameters"></a>

### EthereumKilnStakeParameters
The parameters required for the stake action on Ethereum Kiln.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staker_address | [string](#string) |  | The address you wish to stake from. |
| integrator_contract_address | [string](#string) |  | The address of the integrator contract |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount of Ethereum to stake in wei. |






<a name="coinbase-staking-v1alpha1-EthereumKilnStakingContextDetails"></a>

### EthereumKilnStakingContextDetails
The protocol specific details for an Ethereum Kiln staking context.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ethereum_balance | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The Ethereum balance of the address |
| integrator_share_balance | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The number of integrator shares owned by the address. |
| integrator_share_underlying_balance | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The total Ethereum you can exchange for your integrator shares. |
| total_exitable_eth | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The total amount of Ethereum you can redeem for all non-claimed vPool shares. |
| total_shares_pending_exit | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The number of vPool shares are eligible to receive now or at a later point in time. |
| fulfillable_share_count | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The number of vPool shares you are able to claim now. |






<a name="coinbase-staking-v1alpha1-EthereumKilnStakingContextParameters"></a>

### EthereumKilnStakingContextParameters
The protocol specific parameters required for fetching a staking context.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| integrator_contract_address | [string](#string) |  | Integrator contract address |






<a name="coinbase-staking-v1alpha1-EthereumKilnStakingParameters"></a>

### EthereumKilnStakingParameters
The parameters needed for staking on Ethereum via Kiln.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_parameters | [EthereumKilnStakeParameters](#coinbase-staking-v1alpha1-EthereumKilnStakeParameters) |  | The parameters for stake action on Ethereum Kiln. |
| unstake_parameters | [EthereumKilnUnstakeParameters](#coinbase-staking-v1alpha1-EthereumKilnUnstakeParameters) |  | The parameters for unstake action on Ethereum Kiln. |
| claim_rewards_parameters | [EthereumKilnClaimRewardsParameters](#coinbase-staking-v1alpha1-EthereumKilnClaimRewardsParameters) |  | The parameters for claim rewards action on Ethereum Kiln. |






<a name="coinbase-staking-v1alpha1-EthereumKilnUnstakeParameters"></a>

### EthereumKilnUnstakeParameters
The parameters required for the unstake action on Ethereum Kiln.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staker_address | [string](#string) |  | The address you wish to unstake from. |
| integrator_contract_address | [string](#string) |  | The address of the integrator contract |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount of Ethereum to unstake in wei. |





 

 

 

 



<a name="coinbase_staking_v1alpha1_polygon-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/polygon.proto



<a name="coinbase-staking-v1alpha1-PolygonClaimRewardsParameters"></a>

### PolygonClaimRewardsParameters
The parameters required for claim rewards action on Polygon


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator_address | [string](#string) |  | The public address of the delegator you wish to interact with. |
| validator_address | [string](#string) |  | The public address of the validator you wish to perform the action to. |






<a name="coinbase-staking-v1alpha1-PolygonRestakeParameters"></a>

### PolygonRestakeParameters
The parameters required for unstake action on Polygon


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator_address | [string](#string) |  | The public address of the delegator you wish to interact with. |
| validator_address | [string](#string) |  | The public address of the validator you wish to perform the action to. |






<a name="coinbase-staking-v1alpha1-PolygonStakeParameters"></a>

### PolygonStakeParameters
The parameters required for stake action on Polygon


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator_address | [string](#string) |  | The public address of the delegator you wish to interact with. |
| validator_address | [string](#string) |  | The public address of the validator you wish to perform the action to. |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount of the asset. For native assets or ERC-20 contracts, this is presented in terms of atomic units (e.g., Wei for Ether) as a base-10 number. |






<a name="coinbase-staking-v1alpha1-PolygonStakingParameters"></a>

### PolygonStakingParameters
The parameters needed for staking on Polygon


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_parameters | [PolygonStakeParameters](#coinbase-staking-v1alpha1-PolygonStakeParameters) |  | The parameters for stake action on Polygon |
| unstake_parameters | [PolygonUnstakeParameters](#coinbase-staking-v1alpha1-PolygonUnstakeParameters) |  | The parameters for unstake action on Polygon |
| restake_parameters | [PolygonRestakeParameters](#coinbase-staking-v1alpha1-PolygonRestakeParameters) |  | The parameters for restake action on Polygon |
| claim_rewards_parameters | [PolygonClaimRewardsParameters](#coinbase-staking-v1alpha1-PolygonClaimRewardsParameters) |  | The parameters for claim rewards action on Polygon |






<a name="coinbase-staking-v1alpha1-PolygonUnstakeParameters"></a>

### PolygonUnstakeParameters
The parameters required for unstake action on Polygon


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator_address | [string](#string) |  | The public address of the delegator you wish to interact with. |
| validator_address | [string](#string) |  | The public address of the validator you wish to perform the action to. |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount of the asset. For native assets or ERC-20 contracts, this is presented in terms of atomic units (e.g., Wei for Ether) as a base-10 number. |





 

 

 

 



<a name="coinbase_staking_v1alpha1_solana-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/solana.proto



<a name="coinbase-staking-v1alpha1-NonceOptions"></a>

### NonceOptions
The parameters required to use a durable transaction nonce for Solana transactions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nonce | [string](#string) |  | The blockhash stored in the nonce account. |
| nonce_account | [string](#string) |  | The address of the nonce account. |
| nonce_authority | [string](#string) |  | The address of the nonce authority. If not provided, defaults to the nonce_account_address. Signs the transaction. |






<a name="coinbase-staking-v1alpha1-SolanaCreateStakeAccountParameters"></a>

### SolanaCreateStakeAccountParameters
The parameters required for the create stake account action on Solana.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_account_address | [string](#string) |  | The address of the new stake account which will be created. This address must not already exist. Signs the transaction. |
| from_address | [string](#string) |  | The address of the account which will fund the stake account. Pays the transaction fee. Signs the transaction. |
| stake_authority | [string](#string) |  | The address of the account which will be granted signing authority over staking operations on the new stake account. If not provided, defaults to the from_address. |
| withdraw_authority | [string](#string) |  | The address of the account which will be granted signing authority over withdrawing inactive stake from the new stake account. If not provided, defaults to the from_address. |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount to fund the stake account with, in Lamports. |
| nonce_options | [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions) |  | The parameters required to use a durable transaction nonce for Solana transactions. |






<a name="coinbase-staking-v1alpha1-SolanaDeactivateStakeParameters"></a>

### SolanaDeactivateStakeParameters
The parameters required for the deactivate stake action on Solana.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_account_address | [string](#string) |  | The address of the stake account which will have its stake deactivated. Stake must be currently active. |
| stake_authority | [string](#string) |  | The address of the account which has signing authority over staking operations on the stake account. Signs the transaction. |
| nonce_options | [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions) |  | The parameters required to use a durable transaction nonce for Solana transactions. |






<a name="coinbase-staking-v1alpha1-SolanaDelegateStakeParameters"></a>

### SolanaDelegateStakeParameters
The parameters required for the delegate stake action on Solana.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_account_address | [string](#string) |  | The address of the stake account which will be delegating its stake. |
| vote_account_address | [string](#string) |  | The address of the validator&#39;s vote account to which the stake will be delegated. |
| stake_authority | [string](#string) |  | The address of the account which has signing authority over staking operations on the stake account. Signs the transaction. |
| nonce_options | [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions) |  | The parameters required to use a durable transaction nonce for Solana transactions. |






<a name="coinbase-staking-v1alpha1-SolanaMergeStakeParameters"></a>

### SolanaMergeStakeParameters
The parameters required for merge stake action on Solana.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_account_address | [string](#string) |  | The address of the stake account will be merged into and have stake added to it. |
| source_stake_account_address | [string](#string) |  | The address of the source stake account which will have no longer exist after the merge. |
| stake_authority | [string](#string) |  | The address of the account which has signing authority over staking operations on the stake account. Signs the transaction. |
| nonce_options | [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions) |  | The parameters required to use a durable transaction nonce for Solana transactions. |






<a name="coinbase-staking-v1alpha1-SolanaSplitStakeParameters"></a>

### SolanaSplitStakeParameters
The parameters required for split stake action on Solana.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_account_address | [string](#string) |  | The address of the stake account will be split and have its stake removed. |
| new_stake_account_address | [string](#string) |  | The address of the new stake account which will be created and have the stake added to it. |
| stake_authority | [string](#string) |  | The address of the account which has signing authority over staking operations on the stake account. Signs the transaction. |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount of stake to split, in Lamports. |
| nonce_options | [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions) |  | The parameters required to use a durable transaction nonce for Solana transactions. |






<a name="coinbase-staking-v1alpha1-SolanaStakingParameters"></a>

### SolanaStakingParameters
The parameters needed for staking on Solana


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| create_stake_parameters | [SolanaCreateStakeAccountParameters](#coinbase-staking-v1alpha1-SolanaCreateStakeAccountParameters) |  | The parameters for create stake account action on Solana. |
| delegate_stake_parameters | [SolanaDelegateStakeParameters](#coinbase-staking-v1alpha1-SolanaDelegateStakeParameters) |  | The parameters for delegate stake action on Solana. |
| deactivate_stake_parameters | [SolanaDeactivateStakeParameters](#coinbase-staking-v1alpha1-SolanaDeactivateStakeParameters) |  | The parameters for deactivate stake action on Solana. |
| withdraw_stake_parameters | [SolanaWithdrawStakeParameters](#coinbase-staking-v1alpha1-SolanaWithdrawStakeParameters) |  | The parameters for withdraw stake action on Solana. |
| split_stake_parameters | [SolanaSplitStakeParameters](#coinbase-staking-v1alpha1-SolanaSplitStakeParameters) |  | The parameters for split stake action on Solana. |
| merge_stake_parameters | [SolanaMergeStakeParameters](#coinbase-staking-v1alpha1-SolanaMergeStakeParameters) |  | The parameters for merge stake action on Solana. |






<a name="coinbase-staking-v1alpha1-SolanaWithdrawStakeParameters"></a>

### SolanaWithdrawStakeParameters
The parameters required for the withdraw stake action on Solana.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_account_address | [string](#string) |  | The address of the stake account from which stake will be withdrawn. Stake must be currently inactive. |
| recipient_address | [string](#string) |  | The address of the recipient account which will receive the withdrawn stake. |
| withdraw_authority | [string](#string) |  | The address of the account which has signing authority over withdrawing inactive stake from the stake account. Signs the transaction. |
| amount | [Amount](#coinbase-staking-v1alpha1-Amount) |  | The amount to withdraw from the stake account, in Lamports. |
| nonce_options | [NonceOptions](#coinbase-staking-v1alpha1-NonceOptions) |  | The parameters required to use a durable transaction nonce for Solana transactions. |





 

 

 

 



<a name="coinbase_staking_v1alpha1_workflow-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/workflow.proto



<a name="coinbase-staking-v1alpha1-CreateWorkflowRequest"></a>

### CreateWorkflowRequest
The request message for CreateWorkflow.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parent | [string](#string) |  | The resource name of the parent that owns the workflow. Format: projects/{project} |
| workflow | [Workflow](#coinbase-staking-v1alpha1-Workflow) |  | The workflow to create |






<a name="coinbase-staking-v1alpha1-GetWorkflowRequest"></a>

### GetWorkflowRequest
The message for GetWorkflow.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the workflow. Format: projects/{project}/workflows/{workflow} |






<a name="coinbase-staking-v1alpha1-ListWorkflowsRequest"></a>

### ListWorkflowsRequest
The message for ListWorkflows.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parent | [string](#string) |  | The resource name of the parent that owns the collection of networks. Format: projects/{project} |
| filter | [string](#string) |  | [AIP-160](https://google.aip.dev/160) filter Supported fields: - string delegator_address: &#34;0x...&#34; - string validator_address: &#34;0x...&#34; - string action: &#34;stake&#34;, &#34;unstake&#34; - string protocol: &#34;polygon&#34; - string network: &#34;testnet&#34;, &#34;mainnet&#34; - string amount: &#34;10000&#34; - string currency: &#34;MATIC&#34; |
| page_size | [int32](#int32) |  | The maximum number of workflows to return. The service may return fewer than this value.

If unspecified, 50 workflows will be returned. The maximum value is 1000; values over 1000 will be floored to 1000. |
| page_token | [string](#string) |  | A page token as part of the response of a previous call. Provide this to retrieve the next page.

When paginating, all other parameters must match the previous request to list resources. |






<a name="coinbase-staking-v1alpha1-ListWorkflowsResponse"></a>

### ListWorkflowsResponse
The response message for ListWorkflows.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| workflows | [Workflow](#coinbase-staking-v1alpha1-Workflow) | repeated | The list of workflows. |
| next_page_token | [string](#string) |  | A token which can be provided as `page_token` to retrieve the next page. If this field is omitted, there are no additional pages. |






<a name="coinbase-staking-v1alpha1-PerformWorkflowStepRequest"></a>

### PerformWorkflowStepRequest
The request message for PerformWorkflowStep.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the workflow. Format: projects/{project}/workflows/{workflow} |
| step | [int32](#int32) |  | The index of the step to be performed. |
| data | [string](#string) |  | Transaction metadata. This is either the signed transaction or transaction hash depending on the workflow&#39;s broadcast method. |






<a name="coinbase-staking-v1alpha1-RefreshWorkflowStepRequest"></a>

### RefreshWorkflowStepRequest
The request message for RefreshWorkflowStepRequest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the workflow. Format: projects/{project}/workflows/{workflow} |
| step | [int32](#int32) |  | The index of the step to be refreshed. |






<a name="coinbase-staking-v1alpha1-TxStepOutput"></a>

### TxStepOutput
The details of a transaction being constructed and broadcasted to the network.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unsigned_tx | [string](#string) |  | The unsigned transaction which was signed in order to be broadcasted |
| signed_tx | [string](#string) |  | The signed transaction which was asked to be broadcasted |
| tx_hash | [string](#string) |  | The hash of the broadcasted transaction. |
| state | [TxStepOutput.State](#coinbase-staking-v1alpha1-TxStepOutput-State) |  | The state of the transaction step. |






<a name="coinbase-staking-v1alpha1-WaitStepOutput"></a>

### WaitStepOutput
The output details of a step where we wait for some kind of on-chain activity to finish like reaching a certain checkpoint, epoch or block


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int64](#int64) |  | The beginning of wait period. |
| current | [int64](#int64) |  | The current wait progress. |
| target | [int64](#int64) |  | The target wait end point. |
| unit | [WaitStepOutput.WaitUnit](#coinbase-staking-v1alpha1-WaitStepOutput-WaitUnit) |  | The wait unit (like checkpoint, block, epoch etc) |
| state | [WaitStepOutput.State](#coinbase-staking-v1alpha1-WaitStepOutput-State) |  | The state of the wait step. |






<a name="coinbase-staking-v1alpha1-Workflow"></a>

### Workflow
A Workflow resource


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the workflow. Format: projects/{projectUUID}/workflows/{workflowUUID} Ex: projects/ 123e4567-e89b-12d3-a456-426614174000/workflows/123e4567-e89b-12d3-a456-426614174000 |
| action | [string](#string) |  | The resource name of the action being performed. Format: protocols/{protocol}/networks/{network}/actions/{action} |
| polygon_staking_parameters | [PolygonStakingParameters](#coinbase-staking-v1alpha1-PolygonStakingParameters) |  | Polygon staking parameters |
| solana_staking_parameters | [SolanaStakingParameters](#coinbase-staking-v1alpha1-SolanaStakingParameters) |  | Solana staking parameters |
| ethereum_kiln_staking_parameters | [EthereumKilnStakingParameters](#coinbase-staking-v1alpha1-EthereumKilnStakingParameters) |  | EthereumKiln staking parameters |
| state | [Workflow.State](#coinbase-staking-v1alpha1-Workflow-State) |  | The current state of the workflow |
| current_step_id | [int32](#int32) |  | The index of the current step. |
| steps | [WorkflowStep](#coinbase-staking-v1alpha1-WorkflowStep) | repeated | The list of steps for this workflow. |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp the workflow was created |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp the workflow was last updated |
| skip_broadcast | [bool](#bool) |  | Flag to skip tx broadcast to network on behalf of the user. Use this flag if you instead prefer to broadcast signed txs on your own. |
| complete_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp the workflow completed |






<a name="coinbase-staking-v1alpha1-WorkflowStep"></a>

### WorkflowStep
The information for a step in the workflow


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The human readable name of the step. |
| tx_step_output | [TxStepOutput](#coinbase-staking-v1alpha1-TxStepOutput) |  | The tx step output (i.e. transaction metadata such as unsigned tx, signed tx etc). |
| wait_step_output | [WaitStepOutput](#coinbase-staking-v1alpha1-WaitStepOutput) |  | The waiting details for any kind like how many checkpoints away for unbonding etc |





 


<a name="coinbase-staking-v1alpha1-TxStepOutput-State"></a>

### TxStepOutput.State
State defines an enumeration of states for a staking transaction

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATE_UNSPECIFIED | 0 | Unspecified transaction state, this is for backwards compatibility. |
| STATE_NOT_CONSTRUCTED | 1 | Tx has not yet been constructed in the backend |
| STATE_CONSTRUCTED | 2 | Tx construction is over in the backend |
| STATE_PENDING_SIGNING | 3 | Tx is waiting to be signed |
| STATE_SIGNED | 4 | Tx has been signed and returned to the backend |
| STATE_BROADCASTING | 5 | Tx is being broadcasted to the network |
| STATE_CONFIRMING | 6 | Tx is waiting for confirmation |
| STATE_CONFIRMED | 7 | Tx has been confirmed to be included in a block |
| STATE_FINALIZED | 8 | Tx has been finalized |
| STATE_FAILED | 9 | Tx construction or broadcasting failed |
| STATE_SUCCESS | 10 | Tx has been successfully executed |
| STATE_CANCELING | 11 | Tx is being canceled |
| STATE_CANCELED | 12 | Tx has been canceled |
| STATE_CANCEL_FAILED | 13 | Tx cancellation failed |
| STATE_FAILED_REFRESHABLE | 14 | Tx failed but can be refreshed |
| STATE_REFRESHING | 15 | Tx is being refreshed |
| STATE_PENDING_EXT_BROADCAST | 16 | Tx is waiting to be externally broadcasted by the customer |



<a name="coinbase-staking-v1alpha1-WaitStepOutput-State"></a>

### WaitStepOutput.State
WaitStepState defines an enumeration of states for a wait step

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATE_UNSPECIFIED | 0 | Unspecified wait step state |
| STATE_NOT_STARTED | 1 | Wait step has not started |
| STATE_IN_PROGRESS | 2 | Wait step is in-progress |
| STATE_COMPLETED | 3 | Wait step completed |



<a name="coinbase-staking-v1alpha1-WaitStepOutput-WaitUnit"></a>

### WaitStepOutput.WaitUnit
The unit of wait time

| Name | Number | Description |
| ---- | ------ | ----------- |
| WAIT_UNIT_UNSPECIFIED | 0 | Unspecified wait time |
| WAIT_UNIT_SECONDS | 1 | Wait time measured in seconds |
| WAIT_UNIT_BLOCKS | 2 | Wait time measured in blocks |
| WAIT_UNIT_EPOCHS | 3 | Wait time measured in epochs |
| WAIT_UNIT_CHECKPOINTS | 4 | Wait time measured in checkpoints |



<a name="coinbase-staking-v1alpha1-Workflow-State"></a>

### Workflow.State
The state of a workflow
Example workflow states:
IN_PROGRESS -&gt; WAITING_FOR_SIGNING -&gt; IN_PROGRESS -&gt; COMPLETED
.................................................|-&gt; FAILED
IN_PROGRESS -&gt; WAITING_FOR_SIGNING -&gt; CANCELING -&gt; CANCELED
...............................................|-&gt; CANCEL_FAILED

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATE_UNSPECIFIED | 0 | Unspecified workflow state, this is for backwards compatibility. |
| STATE_IN_PROGRESS | 1 | In Progress represents a workflow that is currently in progress. |
| STATE_WAITING_FOR_SIGNING | 2 | Waiting for signing represents the workflow is waiting on the consumer to sign and return the corresponding signed tx. |
| STATE_COMPLETED | 3 | Completed represents the workflow has completed. |
| STATE_FAILED | 4 | Failed represents the workflow has failed. |
| STATE_CANCELING | 5 | Canceling represents the workflow is being canceled. |
| STATE_CANCELED | 6 | Canceled represents the workflow has been canceled. |
| STATE_CANCEL_FAILED | 7 | Cancel failed represents the workflow failed to cancel. |
| STATE_FAILED_REFRESHABLE | 8 | Failed refreshable represents the workflow failed but can be refreshed. |
| STATE_WAITING_FOR_EXT_BROADCAST | 9 | Waiting for external broadcast represents the workflow is waiting for the customer to broadcast a tx and return its corresponding tx hash. |


 

 

 



<a name="coinbase_staking_v1alpha1_staking_context-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/staking_context.proto



<a name="coinbase-staking-v1alpha1-ViewStakingContextRequest"></a>

### ViewStakingContextRequest
The request message for the ViewStakingContext request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | The address to fetch staking context for. |
| network | [string](#string) |  | The network to fetch staking context for. |
| ethereum_kiln_staking_context_parameters | [EthereumKilnStakingContextParameters](#coinbase-staking-v1alpha1-EthereumKilnStakingContextParameters) |  | EthereumKiln staking context parameters. |






<a name="coinbase-staking-v1alpha1-ViewStakingContextResponse"></a>

### ViewStakingContextResponse
The response message for the ViewStakingContext request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | The address you are getting a staking context for |
| ethereum_kiln_staking_context_details | [EthereumKilnStakingContextDetails](#coinbase-staking-v1alpha1-EthereumKilnStakingContextDetails) |  | EthereumKiln staking context details |





 

 

 

 



<a name="coinbase_staking_v1alpha1_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coinbase/staking/v1alpha1/api.proto


 

 

 


<a name="coinbase-staking-v1alpha1-StakingService"></a>

### StakingService
StakingService manages staking related resources such as protocols, networks, validators and workflows

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProtocols | [ListProtocolsRequest](#coinbase-staking-v1alpha1-ListProtocolsRequest) | [ListProtocolsResponse](#coinbase-staking-v1alpha1-ListProtocolsResponse) | List supported protocols |
| ListNetworks | [ListNetworksRequest](#coinbase-staking-v1alpha1-ListNetworksRequest) | [ListNetworksResponse](#coinbase-staking-v1alpha1-ListNetworksResponse) | List supported staking networks for a given protocol |
| ListValidators | [ListValidatorsRequest](#coinbase-staking-v1alpha1-ListValidatorsRequest) | [ListValidatorsResponse](#coinbase-staking-v1alpha1-ListValidatorsResponse) | List supported validators for a given protocol and network |
| ListActions | [ListActionsRequest](#coinbase-staking-v1alpha1-ListActionsRequest) | [ListActionsResponse](#coinbase-staking-v1alpha1-ListActionsResponse) | List supported actions for a given protocol and network |
| CreateWorkflow | [CreateWorkflowRequest](#coinbase-staking-v1alpha1-CreateWorkflowRequest) | [Workflow](#coinbase-staking-v1alpha1-Workflow) | Create a workflow to perform an action |
| GetWorkflow | [GetWorkflowRequest](#coinbase-staking-v1alpha1-GetWorkflowRequest) | [Workflow](#coinbase-staking-v1alpha1-Workflow) | Get the current state of an active workflow |
| ListWorkflows | [ListWorkflowsRequest](#coinbase-staking-v1alpha1-ListWorkflowsRequest) | [ListWorkflowsResponse](#coinbase-staking-v1alpha1-ListWorkflowsResponse) | List all workflows in a project |
| PerformWorkflowStep | [PerformWorkflowStepRequest](#coinbase-staking-v1alpha1-PerformWorkflowStepRequest) | [Workflow](#coinbase-staking-v1alpha1-Workflow) | Perform the next step in a workflow |
| RefreshWorkflowStep | [RefreshWorkflowStepRequest](#coinbase-staking-v1alpha1-RefreshWorkflowStepRequest) | [Workflow](#coinbase-staking-v1alpha1-Workflow) | Refresh the current step in a workflow |
| ViewStakingContext | [ViewStakingContextRequest](#coinbase-staking-v1alpha1-ViewStakingContextRequest) | [ViewStakingContextResponse](#coinbase-staking-v1alpha1-ViewStakingContextResponse) | View Staking context information given a specific network address |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

