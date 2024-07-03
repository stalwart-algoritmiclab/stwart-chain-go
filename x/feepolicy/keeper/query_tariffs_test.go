/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTariffsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	msgs := createNTariffs(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTariffsRequest
		response *types.QueryGetTariffsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTariffsRequest{
				Denom: msgs[0].Denom,
			},
			response: &types.QueryGetTariffsResponse{Tariffs: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTariffsRequest{
				Denom: msgs[1].Denom,
			},
			response: &types.QueryGetTariffsResponse{Tariffs: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTariffsRequest{
				Denom: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Tariffs(ctx, tc.request)
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

func TestTariffsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	msgs := createNTariffs(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTariffsRequest {
		return &types.QueryAllTariffsRequest{
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
			resp, err := keeper.TariffsAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Tariffs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Tariffs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TariffsAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Tariffs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Tariffs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TariffsAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Tariffs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TariffsAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
