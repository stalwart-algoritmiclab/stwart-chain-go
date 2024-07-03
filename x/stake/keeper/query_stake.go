/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/stake/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakeAll(ctx context.Context, req *types.QueryAllStakeRequest) (*types.QueryAllStakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stakes []types.Stake

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	stakeStore := prefix.NewStore(store, types.KeyPrefix(types.StakeKeyPrefix))

	pageRes, err := query.Paginate(stakeStore, req.Pagination, func(key []byte, value []byte) error {
		var stake types.Stake
		if err := k.cdc.Unmarshal(value, &stake); err != nil {
			return err
		}

		stakes = append(stakes, stake)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStakeResponse{Stake: stakes, Pagination: pageRes}, nil
}

func (k Keeper) Stake(ctx context.Context, req *types.QueryGetStakeRequest) (*types.QueryGetStakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetStake(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStakeResponse{Stake: val}, nil
}
