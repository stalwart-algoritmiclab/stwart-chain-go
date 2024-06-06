/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTariffsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.FeepolicyKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateTariffs{Creator: creator,
			Denom: strconv.Itoa(i),
		}
		_, err := srv.CreateTariffs(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetTariffs(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestTariffsMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateTariffs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateTariffs{Creator: creator,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateTariffs{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateTariffs{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FeepolicyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateTariffs{Creator: creator,
				Denom: strconv.Itoa(0),
			}
			_, err := srv.CreateTariffs(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTariffs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTariffs(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTariffsMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteTariffs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteTariffs{Creator: creator,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteTariffs{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteTariffs{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FeepolicyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateTariffs(ctx, &types.MsgCreateTariffs{Creator: creator,
				Denom: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteTariffs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTariffs(ctx,
					tc.request.Denom,
				)
				require.False(t, found)
			}
		})
	}
}
