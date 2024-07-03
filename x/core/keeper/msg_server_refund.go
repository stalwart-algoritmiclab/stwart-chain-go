/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

func (m msgServer) Refund(goCtx context.Context, msg *types.MsgRefund) (*types.MsgRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: add stake module
	_ = ctx

	return &types.MsgRefundResponse{}, nil
}
