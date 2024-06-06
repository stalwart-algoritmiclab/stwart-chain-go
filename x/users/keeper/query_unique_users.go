/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UniqueUsersAll(
	ctx context.Context,
	req *types.QueryAllUniqueUsersRequest,
) (*types.QueryAllUniqueUsersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var uniqueUsersList []types.UniqueUsers

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	uniqueUsersStore := prefix.NewStore(store, types.KeyPrefix(types.UniqueUsersKeyPrefix))

	pageRes, err := query.Paginate(uniqueUsersStore, req.Pagination, func(key []byte, value []byte) error {
		var uniqueUsers types.UniqueUsers
		if err := k.cdc.Unmarshal(value, &uniqueUsers); err != nil {
			return err
		}

		uniqueUsersList = append(uniqueUsersList, uniqueUsers)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUniqueUsersResponse{UniqueUsers: uniqueUsersList, Pagination: pageRes}, nil
}

func (k Keeper) UniqueUsers(
	ctx context.Context,
	req *types.QueryGetUniqueUsersRequest,
) (*types.QueryGetUniqueUsersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetUniqueUsers(ctx, req.Date)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUniqueUsersResponse{UniqueUsers: val}, nil
}
