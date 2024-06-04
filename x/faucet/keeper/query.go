package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

var _ types.QueryServer = Keeper{}
