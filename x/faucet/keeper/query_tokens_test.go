/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

func TestTokensQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FaucetKeeper(t)
	msgs := createNTokens(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTokensRequest
		response *types.QueryGetTokensResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTokensRequest{Id: msgs[0].Id},
			response: &types.QueryGetTokensResponse{Tokens: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetTokensRequest{Id: msgs[1].Id},
			response: &types.QueryGetTokensResponse{Tokens: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetTokensRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Tokens(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestTokensQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FaucetKeeper(t)
	msgs := createNTokens(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTokensRequest {
		return &types.QueryAllTokensRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TokensAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Tokens), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Tokens),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TokensAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Tokens), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Tokens),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TokensAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Tokens),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TokensAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
