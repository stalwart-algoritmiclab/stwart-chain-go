/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
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
