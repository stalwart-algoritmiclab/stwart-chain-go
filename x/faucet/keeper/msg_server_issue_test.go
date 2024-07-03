/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"testing"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
)

func Test_msgServer_Issue(t *testing.T) {
	k, srv, ctx, accounts := setupMsgServerWithAddresses(t, 2)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
	}

	k.AppendTokens(ctx, types.Tokens{
		Id:      1,
		Denom:   domain.DenomStake,
		Amount:  "100000000",
		Creator: accounts[0].Address.String(),
	})

	type args struct {
		msg *types.MsgIssue
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Completed",
			args: args{
				msg: &types.MsgIssue{
					Creator: accounts[0].Address.String(),
					Address: accounts[1].Address.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "Unauthorized creator",
			args: args{
				msg: &types.MsgIssue{
					Creator: "Some creator",
					Address: accounts[1].Address.String(),
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid address",
			args: args{
				msg: &types.MsgIssue{
					Creator: accounts[0].Address.String(),
					Address: "Some creator",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := srv.Issue(ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Issue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
