/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func (k Keeper) PollsAll(ctx context.Context, req *types.QueryAllPollsRequest) (*types.QueryAllPollsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pollss []types.Polls

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	pollsStore := prefix.NewStore(store, types.KeyPrefix(types.PollsKey))

	pageRes, err := query.Paginate(pollsStore, req.Pagination, func(key []byte, value []byte) error {
		var polls types.Polls
		if err := k.cdc.Unmarshal(value, &polls); err != nil {
			return err
		}

		pollss = append(pollss, polls)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollsResponse{Polls: pollss, Pagination: pageRes}, nil
}

func (k Keeper) Polls(ctx context.Context, req *types.QueryGetPollsRequest) (*types.QueryGetPollsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	polls, found := k.GetPolls(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPollsResponse{Polls: polls}, nil
}
