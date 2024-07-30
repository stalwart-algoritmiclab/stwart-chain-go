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

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func TestVotesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	msgs := createNVotes(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVotesRequest
		response *types.QueryGetVotesResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetVotesRequest{Id: msgs[0].Id},
			response: &types.QueryGetVotesResponse{Votes: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetVotesRequest{Id: msgs[1].Id},
			response: &types.QueryGetVotesResponse{Votes: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetVotesRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Votes(ctx, tc.request)
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

func TestVotesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	msgs := createNVotes(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVotesRequest {
		return &types.QueryAllVotesRequest{
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
			resp, err := keeper.VotesAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Votes), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Votes),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VotesAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Votes), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Votes),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VotesAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Votes),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VotesAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
