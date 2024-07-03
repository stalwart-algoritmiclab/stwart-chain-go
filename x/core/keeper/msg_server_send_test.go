/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"reflect"
	"testing"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
)

func Test_msgServer_Send(t *testing.T) {
	k, ms, ctx, accounts := setupMsgServerWithAddresses(t, 2, domain.DenomStake)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
	}

	// get address from module
	coreAddress, err := k.ModuleAddressesByName(types.ModuleName)
	if err != nil {
		t.Error(err)
	}

	type args struct {
		msg *types.MsgSend
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgSendResponse
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				msg: &types.MsgSend{
					Creator: accounts[0].Address.String(),
					From:    coreAddress,
					To:      accounts[1].Address.String(),
					Amount:  "100000000",
					Denom:   domain.DenomStake,
				},
			},
			want:    &types.MsgSendResponse{},
			wantErr: false,
		},
		{
			name: "Failed empty creator",
			args: args{
				msg: &types.MsgSend{
					Creator: "",
					From:    accounts[0].Address.String(),
					To:      accounts[1].Address.String(),
					Amount:  "100000000",
					Denom:   domain.DenomStake,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty from",
			args: args{
				msg: &types.MsgSend{
					Creator: accounts[0].Address.String(),
					From:    "",
					To:      accounts[1].Address.String(),
					Amount:  "100000000",
					Denom:   domain.DenomStake,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty to",
			args: args{
				msg: &types.MsgSend{
					Creator: accounts[0].Address.String(),
					From:    coreAddress,
					To:      "",
					Amount:  "100000000",
					Denom:   domain.DenomStake,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty amount",
			args: args{
				msg: &types.MsgSend{
					Creator: accounts[0].Address.String(),
					From:    coreAddress,
					To:      accounts[1].Address.String(),
					Amount:  "",
					Denom:   domain.DenomStake,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed zero amount",
			args: args{
				msg: &types.MsgSend{
					Creator: accounts[0].Address.String(),
					From:    coreAddress,
					To:      accounts[1].Address.String(),
					Amount:  "0",
					Denom:   domain.DenomStake,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty denom",
			args: args{
				msg: &types.MsgSend{
					Creator: accounts[0].Address.String(),
					From:    coreAddress,
					To:      accounts[1].Address.String(),
					Amount:  "100000000",
					Denom:   "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ms.Send(ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}
