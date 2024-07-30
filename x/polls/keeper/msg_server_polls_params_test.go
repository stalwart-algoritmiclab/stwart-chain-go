/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func TestPollsParamsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PollsKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	expected := &types.MsgCreatePollsParams{Creator: creator}
	_, err := srv.CreatePollsParams(ctx, expected)
	require.NoError(t, err)
	rst, found := k.GetPollsParams(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestPollsParamsMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePollsParams
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePollsParams{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePollsParams{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PollsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePollsParams{Creator: creator}
			_, err := srv.CreatePollsParams(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePollsParams(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPollsParams(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPollsParamsMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePollsParams
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePollsParams{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePollsParams{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PollsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePollsParams(ctx, &types.MsgCreatePollsParams{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeletePollsParams(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPollsParams(ctx)
				require.False(t, found)
			}
		})
	}
}
