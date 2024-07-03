/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRatesMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RatesKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRates{Creator: creator,
			Denom: strconv.Itoa(i),
		}
		_, err := srv.CreateRates(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRates(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRatesMsgServerUpdate(t *testing.T) {
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"

	tests := []struct {
		desc    string
		request *types.MsgUpdateRates
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRates{
				Creator:  creator,
				Denom:    "usdb",
				Rate:     "1",
				Decimals: 8,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRates{
				Creator:  "B",
				Denom:    "usdtb",
				Rate:     "1",
				Decimals: 8,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRates{
				Creator:  creator,
				Denom:    strconv.Itoa(100000),
				Rate:     "1",
				Decimals: 8,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RatesKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateRates{
				Creator:  creator,
				Denom:    "usdb",
				Rate:     "1",
				Decimals: 8,
			}
			_, err := srv.CreateRates(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRates(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRates(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRatesMsgServerDelete(t *testing.T) {
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"

	tests := []struct {
		desc    string
		request *types.MsgDeleteRates
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRates{Creator: creator,
				Denom: "udtb",
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRates{Creator: "B",
				Denom: "usdb",
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRates{Creator: creator,
				Denom: "usda",
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RatesKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateRates(ctx, &types.MsgCreateRates{
				Creator:  creator,
				Denom:    "udtb",
				Rate:     "1",
				Decimals: 8,
			})
			require.NoError(t, err)
			_, err = srv.DeleteRates(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRates(ctx,
					tc.request.Denom,
				)
				require.False(t, found)
			}
		})
	}
}

func Test_msgServer_CreateRates(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgCreateRates
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgCreateRatesResponse
		wantErr bool
	}{
		{
			name: "",
			args: args{
				goCtx: wctx,
				msg: &types.MsgCreateRates{
					Creator:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:    "usdt",
					Rate:     "1",
					Decimals: 8,
				},
			},
			want:    &types.MsgCreateRatesResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateRates(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_UpdateRates(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgUpdateRates
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgUpdateRatesResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgUpdateRates{
					Creator:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:    "usdc",
					Rate:     "1",
					Decimals: 1,
				},
			},
			want:    &types.MsgUpdateRatesResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.UpdateRates(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_DeleteRates(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)
	type args struct {
		goCtx context.Context
		msg   *types.MsgDeleteRates
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgDeleteRatesResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgDeleteRates{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   "usdc",
				},
			},
			want:    &types.MsgDeleteRatesResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.DeleteRates(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
