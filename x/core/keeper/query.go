/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

var _ types.QueryServer = Keeper{}
