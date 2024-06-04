package ante

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
)

type (
	feesForCoins map[sdkmath.Int]sendFees
	sendFees     struct {
		feeCoin       sdk.Coin
		minRefBalance sdk.Coin
		fee           types.Fees
	}
)
