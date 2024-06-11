/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TariffsAll(ctx context.Context, req *types.QueryAllTariffsRequest) (*types.QueryAllTariffsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tariffss []types.Tariffs

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	tariffsStore := prefix.NewStore(store, types.KeyPrefix(types.TariffsKeyPrefix))

	pageRes, err := query.Paginate(tariffsStore, req.Pagination, func(key []byte, value []byte) error {
		var tariffs types.Tariffs
		if err := k.cdc.Unmarshal(value, &tariffs); err != nil {
			return err
		}

		tariffss = append(tariffss, tariffs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTariffsResponse{Tariffs: tariffss, Pagination: pageRes}, nil
}

func (k Keeper) Tariffs(ctx context.Context, req *types.QueryGetTariffsRequest) (*types.QueryGetTariffsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetTariffs(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTariffsResponse{Tariffs: val}, nil
}
