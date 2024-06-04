package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStatsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SystemrewardsKeeper(t)
	msgs := createNStats(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStatsRequest
		response *types.QueryGetStatsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStatsRequest{
				Date: msgs[0].Date,
			},
			response: &types.QueryGetStatsResponse{Stats: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetStatsRequest{
				Date: msgs[1].Date,
			},
			response: &types.QueryGetStatsResponse{Stats: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetStatsRequest{
				Date: strconv.Itoa(100000),
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
			response, err := keeper.Stats(ctx, tc.request)
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

func TestStatsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SystemrewardsKeeper(t)
	msgs := createNStats(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStatsRequest {
		return &types.QueryAllStatsRequest{
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
			resp, err := keeper.StatsAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Stats), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Stats),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StatsAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Stats), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Stats),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StatsAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Stats),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StatsAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
