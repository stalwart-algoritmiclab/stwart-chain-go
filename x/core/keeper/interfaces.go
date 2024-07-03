/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	CoreKeeper interface {
		BaseKeeper

		AddBurnedToDailyStats(ctx sdk.Context, coins ...sdk.Coin)
		AddIssuedToDailyStats(ctx sdk.Context, coins ...sdk.Coin)
		AddWithdrawnToDailyStats(ctx sdk.Context, coins ...sdk.Coin)
		GetAllStats(ctx context.Context) (list []types.Stats)
		ModulesAddresses(_ context.Context, _ *types.QueryModulesAddressesRequest) (*types.QueryModulesAddressesResponse, error)
		RemoveStats(ctx context.Context, date string)
		StatsAll(ctx context.Context, req *types.QueryAllStatsRequest) (*types.QueryAllStatsResponse, error)
	}

	CoreMsgServer interface {
		CoreKeeper

		Fees(goCtx context.Context, msg *types.MsgFees) (*types.MsgFeesResponse, error)
		Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error)
		RefReward(goCtx context.Context, msg *types.MsgRefReward) (*types.MsgRefRewardResponse, error)
		Refund(goCtx context.Context, msg *types.MsgRefund) (*types.MsgRefundResponse, error)
		Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error)
		UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error)
		Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error)
	}
)
