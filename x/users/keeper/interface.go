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

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) (params types.Params)
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	StatsKeeper interface {
		AddCountUserStats(ctx context.Context, countUsers uint64)
		GetAllStats(ctx context.Context) (list []types.Stats)
		GetStats(ctx context.Context, date string) (val types.Stats, found bool)
		GetStatsByDate(ctx sdk.Context, startDate, endDate string) (result []types.Stats, err error)
		RemoveStats(ctx context.Context, date string)
		SetStats(ctx context.Context, stats types.Stats)
		Stats(ctx context.Context, req *types.QueryGetStatsRequest) (*types.QueryGetStatsResponse, error)
		StatsAll(ctx context.Context, req *types.QueryAllStatsRequest) (*types.QueryAllStatsResponse, error)
		StatsByDate(goCtx context.Context, req *types.QueryStatsByDateRequest) (*types.QueryStatsByDateResponse, error)
	}

	TotalUsersKeeper interface {
		GetTotalUsers(ctx context.Context) types.TotalUsers
		IncrementTotalUsers(ctx context.Context)
		SetTotalUsers(ctx sdk.Context, usersCount uint64)
		Total(goCtx context.Context, req *types.QueryTotalRequest) (*types.QueryTotalResponse, error)
	}

	UniqueUsersKeeper interface {
		GetAllUniqueUsers(ctx context.Context) (list []types.UniqueUsers)
		GetUniqueUsers(ctx context.Context, date string) (val types.UniqueUsers, found bool)
		RemoveUniqueUsers(ctx context.Context, date string)
		SetUniqueUsers(ctx context.Context, uniqueUsers types.UniqueUsers)
		UniqueUsers(ctx context.Context, req *types.QueryGetUniqueUsersRequest) (*types.QueryGetUniqueUsersResponse, error)
		UniqueUsersAll(ctx context.Context, req *types.QueryAllUniqueUsersRequest) (*types.QueryAllUniqueUsersResponse, error)
	}

	UsersKeeper interface {
		BaseKeeper
		StatsKeeper
		TotalUsersKeeper
		UniqueUsersKeeper

		AddNewUserToStat(ctx context.Context)
		CountUsers(ctx context.Context, userAddresses []string) uint64
	}
)
