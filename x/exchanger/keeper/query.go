package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
)

var _ types.QueryServer = Keeper{}
