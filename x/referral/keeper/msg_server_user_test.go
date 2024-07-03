/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUserMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ReferralKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateUser{Creator: creator,
			AccountAddress: strconv.Itoa(i),
		}
		_, err := srv.CreateUser(ctx, expected)
		require.NoError(t, err)
		_, found := k.GetUser(ctx,
			expected.AccountAddress,
		)
		require.True(t, found)
	}
}

func TestUserMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateUser
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateUser{Creator: creator,
				AccountAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateUser{Creator: "B",
				AccountAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateUser{Creator: creator,
				AccountAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ReferralKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateUser{Creator: creator,
				AccountAddress: strconv.Itoa(0),
			}
			_, err := srv.CreateUser(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateUser(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetUser(ctx,
					expected.AccountAddress,
				)
				require.True(t, found)
			}
		})
	}
}

func TestUserMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteUser
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteUser{Creator: creator,
				AccountAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteUser{Creator: "B",
				AccountAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteUser{Creator: creator,
				AccountAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ReferralKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateUser(ctx, &types.MsgCreateUser{Creator: creator,
				AccountAddress: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteUser(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetUser(ctx,
					tc.request.AccountAddress,
				)
				require.False(t, found)
			}
		})
	}
}
