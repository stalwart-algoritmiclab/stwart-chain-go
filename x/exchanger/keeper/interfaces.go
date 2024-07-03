/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	ExchangerKeeper interface {
		BaseKeeper

		Exchange(goCtx context.Context, msg *types.MsgExchange) (*types.MsgExchangeResponse, error)
		UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error)
	}
)
