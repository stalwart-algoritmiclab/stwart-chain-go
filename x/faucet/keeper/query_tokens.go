/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
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

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

func (k Keeper) TokensAll(
	ctx context.Context,
	req *types.QueryAllTokensRequest,
) (*types.QueryAllTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokensList []types.Tokens

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	tokensStore := prefix.NewStore(store, types.KeyPrefix(types.TokensKey))

	pageRes, err := query.Paginate(tokensStore, req.Pagination, func(key []byte, value []byte) error {
		var tokens types.Tokens
		if err := k.cdc.Unmarshal(value, &tokens); err != nil {
			return err
		}

		tokensList = append(tokensList, tokens)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokensResponse{Tokens: tokensList, Pagination: pageRes}, nil
}

func (k Keeper) Tokens(ctx context.Context, req *types.QueryGetTokensRequest) (*types.QueryGetTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	tokens, found := k.GetTokens(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTokensResponse{Tokens: tokens}, nil
}
