package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stake/types"
)

var _ types.QueryServer = Keeper{}
