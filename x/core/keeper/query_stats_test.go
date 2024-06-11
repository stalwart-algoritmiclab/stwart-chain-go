/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	d "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStatsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
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
	keeper, ctx := keepertest.CoreKeeper(t)
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

func TestKeeper_StatsByDate(t *testing.T) {
	k, _, goCtx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create some stats. Note: used ordered denoms
	denoms := []string{d.DenomStableIndex, d.DenomStake}
	coins := make(sdk.Coins, 0, len(denoms))
	for _, denom := range denoms {
		coins = append(coins, sdk.NewInt64Coin(denom, 1_0000_0000))
	}

	count := uint64(len(coins))

	k.AddBurnedToDailyStats(ctx, coins...)
	k.AddIssuedToDailyStats(ctx, coins...)
	k.AddWithdrawnToDailyStats(ctx, coins...)

	yesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
	today := time.Now().Format(time.DateOnly)

	type args struct {
		req *types.QueryGetStatsByDateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryAllStatsResponse
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				req: &types.QueryGetStatsByDateRequest{
					StartDate: yesterday,
					EndDate:   today,
					Pagination: &query.PageRequest{
						Offset:     0,
						Limit:      1,
						CountTotal: false,
						Reverse:    false,
					},
				},
			},
			want: &types.QueryAllStatsResponse{
				Stats: []types.Stats{
					{
						Date: today,
						DailyStats: &types.DailyStats{
							IssuedCoins:   coins,
							CountIssued:   count,
							BurnedCoins:   coins,
							CountBurned:   count,
							WithdrawCoins: coins,
							CountWithdraw: count,
						},
					},
				},
				Pagination: &query.PageResponse{NextKey: []byte("p_core")},
			},
			wantErr: false,
		},
		{
			name: "Success no data",
			args: args{
				req: &types.QueryGetStatsByDateRequest{
					StartDate: yesterday,
					EndDate:   yesterday,
					Pagination: &query.PageRequest{
						Offset:     0,
						Limit:      1,
						CountTotal: false,
						Reverse:    false,
					},
				},
			},
			want: &types.QueryAllStatsResponse{
				Stats:      []types.Stats{},
				Pagination: &query.PageResponse{NextKey: []byte("p_core")},
			},
			wantErr: false,
		},
		{
			name: "Failed nil request",
			args: args{
				req: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty start date",
			args: args{
				req: &types.QueryGetStatsByDateRequest{
					StartDate: "",
					EndDate:   today,
					Pagination: &query.PageRequest{
						Offset:     0,
						Limit:      1,
						CountTotal: false,
						Reverse:    false,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty end date",
			args: args{
				req: &types.QueryGetStatsByDateRequest{
					StartDate: yesterday,
					EndDate:   "",
					Pagination: &query.PageRequest{
						Offset:     0,
						Limit:      1,
						CountTotal: false,
						Reverse:    false,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed start date after end date",
			args: args{
				req: &types.QueryGetStatsByDateRequest{
					StartDate: today,
					EndDate:   yesterday,
					Pagination: &query.PageRequest{
						Offset:     0,
						Limit:      1,
						CountTotal: false,
						Reverse:    false,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed nil pagination",
			args: args{
				req: &types.QueryGetStatsByDateRequest{
					StartDate:  yesterday,
					EndDate:    today,
					Pagination: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := k.StatsByDate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatsByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatsByDate()\ngot = %v,\nwnt = %v", got, tt.want)
			}
		})
	}
}
