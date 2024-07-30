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

func (k Keeper) OptionsAll(ctx context.Context, req *types.QueryAllOptionsRequest) (*types.QueryAllOptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var optionss []types.Options

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	optionsStore := prefix.NewStore(store, types.KeyPrefix(types.OptionsKey))

	pageRes, err := query.Paginate(optionsStore, req.Pagination, func(key []byte, value []byte) error {
		var options types.Options
		if err := k.cdc.Unmarshal(value, &options); err != nil {
			return err
		}

		optionss = append(optionss, options)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOptionsResponse{Options: optionss, Pagination: pageRes}, nil
}

func (k Keeper) Options(ctx context.Context, req *types.QueryGetOptionsRequest) (*types.QueryGetOptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	options, found := k.GetOptions(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOptionsResponse{Options: options}, nil
}
