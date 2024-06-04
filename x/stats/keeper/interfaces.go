package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"
)

type (
	FeeStats interface {
		GetAllFeeStats(ctx context.Context) (list []types.FeeStats)
		GetOneStatsByDate(ctx sdk.Context, date string) (stat types.FeeStats, found bool)
		GetFeeStats(ctx context.Context, index string) (val types.FeeStats, found bool)
		GetStatsByDate(ctx sdk.Context, startDate, endDate string) (list []types.FeeStats, err error)
		GetStatsByIndexes(ctx sdk.Context, startIndex, endIndex string) (list []types.FeeStats, found bool)
		RemoveFeeStats(ctx context.Context, index string)
		SetFeeStats(ctx context.Context, stats types.FeeStats)
		SetStatsFee(ctx sdk.Context, amountFee sdk.Coins, amountTx sdk.Coins)
		SetStatsNoFee(ctx sdk.Context, amount sdk.Coins)
		FeeStatsAll(goCtx context.Context, req *types.QueryAllFeeStatsRequest) (*types.QueryAllFeeStatsResponse, error)
		StatsByDate(goCtx context.Context, req *types.QueryGetFeeStatsByDateRequest) (*types.QueryAllFeeStatsResponse, error)
		StatsByIndexes(goCtx context.Context, req *types.QueryGetFeeStatsByIndexesRequest) (*types.QueryAllFeeStatsResponse, error)
	}
)
