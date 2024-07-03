/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	core "github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	rewards "github.com/stalwart-algoritmiclab/stwart-chain-go/x/systemrewards/types"
	user "github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

type (
	CoreKeeper interface {
		GetStatsByDate(ctx sdk.Context, startDate, endDate string) (result []core.Stats, err error)
	}

	RewardsKeeper interface {
		GetStatsByDate(ctx context.Context, startDate, endDate string) (result []rewards.Stats, err error)
	}

	UserKeeper interface {
		GetStatsByDate(ctx sdk.Context, startDate, endDate string) (result []user.Stats, err error)
	}
)
