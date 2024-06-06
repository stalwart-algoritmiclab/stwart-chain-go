/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUser(keeper keeper.Keeper, ctx context.Context, n int) []types.User {
	items := make([]types.User, n)
	for i := range items {
		items[i].AccountAddress = strconv.Itoa(i)

		keeper.SetUser(ctx, items[i])
	}
	return items
}

func TestUserGet(t *testing.T) {
	keeper, ctx := keepertest.ReferralKeeper(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUser(ctx,
			item.AccountAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserRemove(t *testing.T) {
	keeper, ctx := keepertest.ReferralKeeper(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUser(ctx,
			item.AccountAddress,
		)
		_, found := keeper.GetUser(ctx,
			item.AccountAddress,
		)
		require.False(t, found)
	}
}

func TestUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.ReferralKeeper(t)
	items := createNUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUser(ctx)),
	)
}
