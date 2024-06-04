package keeper_test

import (
	"context"
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"github.com/stretchr/testify/require"
)

func createNAddresses(keeper keeper.Keeper, ctx context.Context, n int) []types.Address {
	items := make([]types.Address, n)
	for i := range items {
		items[i].Id = keeper.AppendAddresses(ctx, items[i])
	}
	return items
}

func TestAddressesGet(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAddresses(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAddressesRemove(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddresses(ctx, item.Id)
		_, found := keeper.GetAddresses(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAddressesGetAll(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAddresses(ctx)),
	)
}

func TestAddressesCount(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAddressesCount(ctx))
}
