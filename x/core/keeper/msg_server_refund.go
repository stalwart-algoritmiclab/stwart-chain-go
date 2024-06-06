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

func (k msgServer) Refund(goCtx context.Context, msg *types.MsgRefund) (*types.MsgRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: add stake module
	_ = ctx

	return &types.MsgRefundResponse{}, nil
}
