package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTariff(keeper keeper.Keeper, ctx context.Context, n int) []types.Tariff {
	items := make([]types.Tariff, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetTariff(ctx, items[i])
	}
	return items
}

func TestTariffGet(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariff(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTariff(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTariffRemove(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariff(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTariff(ctx,
			item.Denom,
		)
		_, found := keeper.GetTariff(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestTariffGetAll(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariff(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTariff(ctx)),
	)
}
