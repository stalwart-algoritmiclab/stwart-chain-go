/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/rand"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/keeper"
	securedtypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"

	"github.com/stretchr/testify/require"
)

func createNAddresses(keeper keeper.Keeper, ctx context.Context, n int) []securedtypes.Addresses {
	items := make([]securedtypes.Addresses, n)
	for i := range items {
		items[i].Id = keeper.AppendAddresses(ctx, items[i])
	}
	return items
}

func TestAddressesGet(t *testing.T) {
	keeper, ctx := keepertest.SecuredKeeper(t)
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
	keeper, ctx := keepertest.SecuredKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveByID(ctx, item.Id)
		_, found := keeper.GetAddresses(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAddressesGetAll(t *testing.T) {
	keeper, ctx := keepertest.SecuredKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAddresses(ctx)),
	)
}

func TestAddressesCount(t *testing.T) {
	keeper, ctx := keepertest.SecuredKeeper(t)
	items := createNAddresses(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAddressesCount(ctx))
}

func TestKeeper_GetAddressesByAddress(t *testing.T) {
	k, _, goCtx := setupMsgServer(t)
	account, _ := simtypes.RandomAcc(rand.NewRand(), simtypes.RandomAccounts(rand.NewRand(), 1))
	k.SetAddresses(goCtx, securedtypes.Addresses{
		Id: 1,
		Address: []string{
			account.Address.String(),
		},
		Creator: account.Address.String(),
	})

	type args struct {
		address string
	}
	tests := []struct {
		name      string
		args      args
		wantFound bool
	}{
		{
			name: "Success",
			args: args{
				address: account.Address.String(),
			},
			wantFound: true,
		},
		{
			name: "Not found",
			args: args{
				address: "Some address",
			},
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, ok := k.GetAddressesByAddress(types.UnwrapSDKContext(goCtx), tt.args.address); ok != tt.wantFound {
				t.Errorf("GetAddressesByAddress() gotFound = %v, want %v", ok, tt.wantFound)
			}
		})
	}
}
