/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
	staketypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stake/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	StakeKeeper interface {
		SetStake(ctx context.Context, stake staketypes.Stake)
		GetStake(ctx context.Context, address string) (val staketypes.Stake, found bool)
		GetFreeStake(ctx context.Context, address sdk.AccAddress) sdkmath.Int
		GetAllStake(ctx context.Context) (list []staketypes.Stake)
		RemoveStake(ctx context.Context, address string)
	}
)
