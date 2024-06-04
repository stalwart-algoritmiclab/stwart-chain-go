package keeper

import (
	"context"
	"time"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StatsByDate - returns stats by date
func (k Keeper) StatsByDate(goCtx context.Context, req *types.QueryStatsByDateRequest) (*types.QueryStatsByDateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	startDate, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid startDate")
	}

	endDate, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid endDate")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	result := make([]types.Stats, 0)

	paginateResp, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var val types.Stats
		if err = k.cdc.Unmarshal(value, &val); err != nil {
			return status.Error(codes.Internal, "invalid stored stats")
		}

		dateTime, err := time.Parse(time.DateOnly, val.Date)
		if err != nil {
			return status.Error(codes.Internal, "invalid stored field date")
		}

		if dateTime.Before(startDate) || dateTime.After(endDate) {
			return nil
		}

		result = append(result, val)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryStatsByDateResponse{Stats: result, Pagination: paginateResp}, nil
}
