/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"reflect"
	"testing"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/rand"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	d "github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
)

func Test_msgServer_Withdraw(t *testing.T) {
	_, ms, ctx, accounts := setupMsgServerWithAddresses(t, 2, d.DenomStableIndex, d.DenomStake)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
	}

	type args struct {
		msg *types.MsgWithdraw
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgWithdrawResponse
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   d.DenomStableIndex,
					Address: accounts[1].Address.String(),
				},
			},
			want:    &types.MsgWithdrawResponse{},
			wantErr: false,
		},
		{
			name: "Failed withdraw stake",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   d.DenomStake,
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty creator",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: "",
					Amount:  "100000000",
					Denom:   d.DenomStableIndex,
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty amount",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "",
					Denom:   d.DenomStableIndex,
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed zero amount",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "0",
					Denom:   d.DenomStableIndex,
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty denom",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   "",
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty address",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   d.DenomStableIndex,
					Address: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed address not found",
			args: args{
				msg: &types.MsgWithdraw{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   d.DenomStableIndex,
					Address: simtypes.RandomAccounts(rand.NewRand(), 1)[0].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ms.Withdraw(ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Withdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}
