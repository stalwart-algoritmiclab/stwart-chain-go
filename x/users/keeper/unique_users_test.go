package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUniqueUsers(keeper keeper.Keeper, ctx context.Context, n int) []types.UniqueUsers {
	items := make([]types.UniqueUsers, n)
	for i := range items {
		items[i].Date = strconv.Itoa(i)

		keeper.SetUniqueUsers(ctx, items[i])
	}
	return items
}

func TestUniqueUsersGet(t *testing.T) {
	keeper, ctx := keepertest.UsersKeeper(t)
	items := createNUniqueUsers(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniqueUsers(ctx,
			item.Date,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUniqueUsersRemove(t *testing.T) {
	keeper, ctx := keepertest.UsersKeeper(t)
	items := createNUniqueUsers(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniqueUsers(ctx,
			item.Date,
		)
		_, found := keeper.GetUniqueUsers(ctx,
			item.Date,
		)
		require.False(t, found)
	}
}

func TestUniqueUsersGetAll(t *testing.T) {
	keeper, ctx := keepertest.UsersKeeper(t)
	items := createNUniqueUsers(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniqueUsers(ctx)),
	)
}
