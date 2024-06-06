/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

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
