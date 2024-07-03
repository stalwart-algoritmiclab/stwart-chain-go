/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
)

func TestAddressesMsgServerExchange(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	_, err := srv.Exchange(wctx, &types.MsgExchange{
		Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
		Denom:   domain.DenomStableIndex,
		Amount:  "10000000",
		DenomTo: "usdt",
	})
	require.NoError(t, err)

}

func Test_msgServer_Exchange(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgExchange
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgExchangeResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS] ssc to usdt",
			args: args{
				goCtx: wctx,
				msg: &types.MsgExchange{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   domain.DenomStableIndex,
					Amount:  "10000000",
					DenomTo: "usdt",
				},
			},
			want:    &types.MsgExchangeResponse{},
			wantErr: false,
		},
		{
			name: "[SUCCESS] usdt to ssc",
			args: args{
				goCtx: wctx,
				msg: &types.MsgExchange{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   "usdt",
					Amount:  "10000000",
					DenomTo: domain.DenomStableIndex,
				},
			},
			want:    &types.MsgExchangeResponse{},
			wantErr: false,
		},
		{
			name: "[FAIL] small amount ssc to usdt",
			args: args{
				goCtx: wctx,
				msg: &types.MsgExchange{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   domain.DenomStableIndex,
					Amount:  "100",
					DenomTo: "usdt",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[FAIL] small amount usdt to ssc ",
			args: args{
				goCtx: wctx,
				msg: &types.MsgExchange{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   "usdt",
					Amount:  "100",
					DenomTo: domain.DenomStableIndex,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[FAIL] insufficient funds to exchange usdt to ssc",
			args: args{
				goCtx: wctx,
				msg: &types.MsgExchange{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   "usdt",
					Amount:  "10000000000",
					DenomTo: domain.DenomStableIndex,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[FAIL] not found rate",
			args: args{
				goCtx: wctx,
				msg: &types.MsgExchange{
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   "usdc",
					Amount:  "100000",
					DenomTo: domain.DenomStableIndex,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.Exchange(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exchange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exchange() got = %v, want %v", got, tt.want)
			}
		})
	}
}
