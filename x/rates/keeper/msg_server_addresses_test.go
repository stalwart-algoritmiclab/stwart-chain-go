/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"context"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

func TestAddressesMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAddresses(wctx, &types.MsgCreateAddresses{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAddressesMsgServerUpdate(t *testing.T) {
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"

	tests := []struct {
		desc    string
		request *types.MsgUpdateAddresses
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAddresses{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAddresses{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAddresses{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateAddresses(wctx, &types.MsgCreateAddresses{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAddresses(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddressesMsgServerDelete(t *testing.T) {
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"

	tests := []struct {
		desc    string
		request *types.MsgDeleteAddresses
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAddresses{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAddresses{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAddresses{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateAddresses(wctx, &types.MsgCreateAddresses{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAddresses(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_msgServer_CreateAddresses(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgCreateAddresses
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgCreateAddressesResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgCreateAddresses{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Address: []string{"stwart1r92s60x75ashfyjy034kxzdml43y6xrqcafk69"},
				},
			},
			want: &types.MsgCreateAddressesResponse{
				Id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateAddresses(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAddresses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_UpdateAddresses(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgUpdateAddresses
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgUpdateAddressesResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgUpdateAddresses{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Id:      0,
					Address: []string{"stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"},
				},
			},
			want:    &types.MsgUpdateAddressesResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.UpdateAddresses(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAddresses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_DeleteAddresses(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgDeleteAddresses
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgDeleteAddressesResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgDeleteAddresses{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Id:      0,
				},
			},
			want:    &types.MsgDeleteAddressesResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.DeleteAddresses(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAddresses() got = %v, want %v", got, tt.want)
			}
		})
	}
}
