/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	d "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStats(keeper keeper.Keeper, ctx context.Context, n int) []types.Stats {
	items := make([]types.Stats, n)
	for i := range items {
		items[i].Date = strconv.Itoa(i)
		keeper.SetStats(ctx, items[i])
	}
	return items
}

func TestStatsGet(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	items := createNStats(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStats(ctx, item.Date)
		require.True(t, found)
		require.Equal(t, nullify.Fill(&item), nullify.Fill(&rst))
	}
}
func TestStatsRemove(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	items := createNStats(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStats(ctx, item.Date)
		_, found := keeper.GetStats(ctx, item.Date)
		require.False(t, found)
	}
}

func TestStatsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	items := createNStats(keeper, ctx, 10)
	require.ElementsMatch(t, nullify.Fill(items), nullify.Fill(keeper.GetAllStats(ctx)))
}

func TestKeeper_GetStatsByDate(t *testing.T) {
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

	today := time.Now().Format(time.DateOnly)
	tomorrow := time.Now().AddDate(0, 0, 1).Format(time.DateOnly)

	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []types.Stats
		wantErr    bool
	}{
		{
			name: "Success",
			args: args{
				startDate: today,
				endDate:   tomorrow,
			},
			wantResult: []types.Stats{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := k.GetStatsByDate(ctx, tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatsByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetStatsByDate() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
