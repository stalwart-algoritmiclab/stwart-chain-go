/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"reflect"
	"testing"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"
)

func Test_msgServer_Issue(t *testing.T) {
	_, ms, ctx, accounts := setupMsgServerWithAddresses(t, 2)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
	}

	type args struct {
		msg *types.MsgIssue
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgIssueResponse
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				msg: &types.MsgIssue{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   domain.DenomStake,
					Address: accounts[1].Address.String(),
				},
			},
			want:    &types.MsgIssueResponse{},
			wantErr: false,
		},
		{
			name: "Failed empty creator",
			args: args{
				msg: &types.MsgIssue{
					Creator: "",
					Amount:  "100000000",
					Denom:   domain.DenomStake,
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty address",
			args: args{
				msg: &types.MsgIssue{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   domain.DenomStake,
					Address: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty amount",
			args: args{
				msg: &types.MsgIssue{
					Creator: accounts[0].Address.String(),
					Amount:  "",
					Denom:   domain.DenomStake,
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed empty denom",
			args: args{
				msg: &types.MsgIssue{
					Creator: accounts[0].Address.String(),
					Amount:  "100000000",
					Denom:   "",
					Address: accounts[1].Address.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ms.Issue(ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Issue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
