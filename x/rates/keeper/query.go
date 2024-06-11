/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

var _ types.QueryServer = Keeper{}
