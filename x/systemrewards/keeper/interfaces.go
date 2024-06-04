package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"
)

type (
	BaseKeeper interface {
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	// SystemRewardsKeeper used as holder for all system rewards
	SystemRewardsKeeper interface {
		AddStats(ctx context.Context, coins ...sdk.Coin)
		GetStatsByDate(ctx context.Context, startDate, endDate string) (result []types.Stats, err error)
		GetAllStats(ctx context.Context) (list []types.Stats)
		GetStats(ctx context.Context, date string) (val types.Stats, found bool)
		RemoveStats(ctx context.Context, date string)
		SetStats(goCtx context.Context, stats types.Stats)
		Stats(goCtx context.Context, req *types.QueryGetStatsRequest) (*types.QueryGetStatsResponse, error)
		StatsAll(goCtx context.Context, req *types.QueryAllStatsRequest) (*types.QueryAllStatsResponse, error)
		StatsByDate(goCtx context.Context, req *types.QueryStatsByDateRequest) (*types.QueryStatsByDateResponse, error)
	}
)
