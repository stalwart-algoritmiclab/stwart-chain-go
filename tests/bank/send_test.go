/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package bank_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"
)

func TestSend(t *testing.T) {
	amount := math.NewInt(10000)
	bk, ctx, accounts := keeper.BankKeeperWithAccounts(t, 2, domain.DenomStableIndex, domain.DenomStake)

	type args struct {
		fromAddr types.AccAddress
		toAddr   types.AccAddress
		amt      types.Coins
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success send coin: stake",
			args: args{
				fromAddr: accounts[0].Address,
				toAddr:   accounts[1].Address,
				amt:      types.NewCoins(types.NewCoin(domain.DenomStake, amount)),
			},
			wantErr: false,
		},
		{
			name: "Success send coin: stable index",
			args: args{
				fromAddr: accounts[0].Address,
				toAddr:   accounts[1].Address,
				amt:      types.NewCoins(types.NewCoin(domain.DenomStableIndex, amount)),
			},
			wantErr: false,
		},
		{
			name: "Success send all coins",
			args: args{
				fromAddr: accounts[0].Address,
				toAddr:   accounts[1].Address,
				amt: types.NewCoins(
					types.NewCoin(domain.DenomStableIndex, amount),
					types.NewCoin(domain.DenomStake, amount),
				),
			},
			wantErr: false,
		},
		{
			name: "Success empty coins",
			args: args{
				fromAddr: accounts[0].Address,
				toAddr:   accounts[1].Address,
				amt:      nil,
			},
			wantErr: false,
		},
		{
			name: "Success empty to address",
			args: args{
				fromAddr: accounts[0].Address,
				toAddr:   nil,
				amt:      types.NewCoins(types.NewCoin(domain.DenomStake, amount)),
			},
			wantErr: false,
		},
		{
			name: "Failed non-existent denom",
			args: args{
				fromAddr: accounts[0].Address,
				toAddr:   accounts[1].Address,
				amt:      types.NewCoins(types.NewCoin("test", amount)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coinsFrom := bk.SpendableCoins(ctx, tt.args.fromAddr)
			coinsTo := bk.SpendableCoins(ctx, tt.args.toAddr)

			if err := bk.SendCoins(ctx, tt.args.fromAddr, tt.args.toAddr, tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("SendCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("acc from: %v -> %v", coinsFrom, bk.SpendableCoins(ctx, tt.args.fromAddr))
			t.Logf("acc to: %v -> %v", coinsTo, bk.SpendableCoins(ctx, tt.args.toAddr))
		})
	}
}
