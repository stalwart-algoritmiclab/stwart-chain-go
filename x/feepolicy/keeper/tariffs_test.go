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

func createNTariffs(keeper keeper.Keeper, ctx context.Context, n int) []types.Tariffs {
	items := make([]types.Tariffs, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetTariffs(ctx, items[i])
	}
	return items
}

func TestTariffsGet(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariffs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTariffs(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTariffsRemove(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariffs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTariffs(ctx,
			item.Denom,
		)
		_, found := keeper.GetTariffs(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestTariffsGetAll(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariffs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTariffs(ctx)),
	)
}