/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.AnteDecorator = CoreDecorator{}

// CoreDecorator consume some gas amount depending on send amount and calls next ante handle with an updated context
type CoreDecorator struct {
	ak          AccountKeeper
	bk          BankKeeper
	fek         FeePolicyKeeper
	corek       CoreKeeper
	refk        RefKeeper
	rewk        RewardsKeeper
	uk          UserKeeper
	stakeKeeper StakeKeeper
	statsKeeper StatsKeeper
}

// NewCoreDecorator takes an BankingKeeper and returns a new fee consumer for send transaction
func NewCoreDecorator(
	ak AccountKeeper,
	bk BankKeeper,
	fek FeePolicyKeeper,
	corek CoreKeeper,
	refk RefKeeper,
	uk UserKeeper,
	rewk RewardsKeeper,
	stakeKeeper StakeKeeper,
	statsKeeper StatsKeeper,
) CoreDecorator {
	return CoreDecorator{
		ak:          ak,
		bk:          bk,
		fek:         fek,
		corek:       corek,
		refk:        refk,
		rewk:        rewk,
		uk:          uk,
		stakeKeeper: stakeKeeper,
		statsKeeper: statsKeeper,
	}
}
