/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	FaucetKeeper interface {
		BaseKeeper

		CreateTokens(goCtx context.Context, msg *types.MsgCreateTokens) (*types.MsgCreateTokensResponse, error)
		DeleteTokens(goCtx context.Context, msg *types.MsgDeleteTokens) (*types.MsgDeleteTokensResponse, error)
		Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error)
		UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error)
		UpdateTokens(goCtx context.Context, msg *types.MsgUpdateTokens) (*types.MsgUpdateTokensResponse, error)
	}
)
