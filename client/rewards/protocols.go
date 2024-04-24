package rewards

import (
	rewardspb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
)

var (
	Ethereum = rewardspb.ProtocolResourceName{Protocol: "ethereum"}.String()
	Solana   = rewardspb.ProtocolResourceName{Protocol: "solana"}.String()
	Cosmos   = rewardspb.ProtocolResourceName{Protocol: "cosmos"}.String()
)
