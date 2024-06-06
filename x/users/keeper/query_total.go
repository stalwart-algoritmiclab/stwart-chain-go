/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"
	"fmt"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Total - returns a count of users
func (k Keeper) Total(goCtx context.Context, req *types.QueryTotalRequest) (*types.QueryTotalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryTotalResponse{Count: fmt.Sprintf("%d", k.GetTotalUsers(ctx).Total)}, nil
}
