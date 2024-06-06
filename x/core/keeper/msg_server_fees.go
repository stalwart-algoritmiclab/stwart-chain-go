/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Fees(goCtx context.Context, msg *types.MsgFees) (*types.MsgFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgFeesResponse{}, nil
}
