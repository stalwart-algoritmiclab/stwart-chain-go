/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"testing"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

func TestKeeper_ModulesAddresses(t *testing.T) {
	k, _, ctx, accounts := setupMsgServerWithAddresses(t, 2)
	if len(accounts) < 2 {
		t.Error("must have at least 2 accounts")
	}

	type args struct {
		req *types.QueryModulesAddressesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{req: &types.QueryModulesAddressesRequest{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := k.ModulesAddresses(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModulesAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("got: %v", got)
		})
	}
}
