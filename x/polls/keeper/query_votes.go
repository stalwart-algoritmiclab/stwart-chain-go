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

func (k Keeper) VotesAll(ctx context.Context, req *types.QueryAllVotesRequest) (*types.QueryAllVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votess []types.Votes

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	votesStore := prefix.NewStore(store, types.KeyPrefix(types.VotesKey))

	pageRes, err := query.Paginate(votesStore, req.Pagination, func(key []byte, value []byte) error {
		var votes types.Votes
		if err := k.cdc.Unmarshal(value, &votes); err != nil {
			return err
		}

		votess = append(votess, votes)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVotesResponse{Votes: votess, Pagination: pageRes}, nil
}

func (k Keeper) Votes(ctx context.Context, req *types.QueryGetVotesRequest) (*types.QueryGetVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	votes, found := k.GetVotes(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVotesResponse{Votes: votes}, nil
}
