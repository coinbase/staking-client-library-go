package rewards

import (
	api "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
)

var (
	Ethereum = api.ProtocolResourceName{Protocol: "ethereum"}.String()
	Solana   = api.ProtocolResourceName{Protocol: "solana"}.String()
	Cosmos   = api.ProtocolResourceName{Protocol: "cosmos"}.String()
)
