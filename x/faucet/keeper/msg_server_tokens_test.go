/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/rand"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

func TestTokensMsgServerCreate(t *testing.T) {
	_, srv, ctx, accounts := setupMsgServerWithAddresses(t, 3)

	wctx := sdk.UnwrapSDKContext(ctx)
	simAccount, _ := simtypes.RandomAcc(rand.NewRand(), accounts)

	creator := simAccount.Address.String()

	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTokens(wctx, &types.MsgCreateTokens{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTokensMsgServerUpdate(t *testing.T) {
	_, srv, ctx, accounts := setupMsgServerWithAddresses(t, 2)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
		return
	}

	creator := accounts[0].Address.String()

	tests := []struct {
		desc    string
		request *types.MsgUpdateTokens
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTokens{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTokens{Creator: "Some creator"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTokens{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateTokens(wctx, &types.MsgCreateTokens{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTokens(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTokensMsgServerDelete(t *testing.T) {
	_, srv, ctx, accounts := setupMsgServerWithAddresses(t, 2)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
		return
	}

	creator := accounts[0].Address.String()

	tests := []struct {
		desc    string
		request *types.MsgDeleteTokens
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTokens{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTokens{Creator: "Some creator"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTokens{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateTokens(wctx, &types.MsgCreateTokens{Creator: creator})
			require.NoError(t, err)

			if _, err = srv.DeleteTokens(wctx, tc.request); tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
