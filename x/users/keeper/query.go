package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"
)

var _ types.QueryServer = Keeper{}
