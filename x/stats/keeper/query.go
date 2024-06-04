package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"
)

var _ types.QueryServer = Keeper{}
